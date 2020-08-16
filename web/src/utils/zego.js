import { ZegoExpressEngine } from 'zego-express-engine-webrtc';

const userID = 'sample' + new Date().getTime();
const userName = 'sampleUser' + new Date().getTime();
const tokenUrl = 'https://wsliveroom-demo.zego.im:8282/token';
const publishStreamId = 'webrtc' + new Date().getTime();
let zg;
let appID = 1739272706;
let server = 'wss://webliveroom-test.zego.im/ws'; //'wss://wsliveroom' + appID + '-api.zego.im:8282/ws'
let previewVideo;
let useLocalStreamList = [];
let isPreviewed = false;
let supportScreenSharing = false;
let loginRoom = false;

let localStream;
let publishType;

zg = new ZegoExpressEngine(appID, server);
window.zg = zg;

async function checkAnRun(checkScreen) {
    console.log('sdk version is', zg.getVersion());
    const result = await zg.checkSystemRequirements();

    console.warn('checkSystemRequirements ', result);
    // !result.videoCodec.H264 && $('#videoCodeType option:eq(1)').attr('disabled', 'disabled');
    // !result.videoCodec.VP8 && $('#videoCodeType option:eq(2)').attr('disabled', 'disabled');

    if (!result.webRTC) {
        alert('browser is not support webrtc!!');
        return false;
    } else if (!result.videoCodec.H264 && !result.videoCodec.VP8) {
        alert('browser is not support H264 and VP8');
        return false;
    } else if (result.videoCodec.H264) {
        supportScreenSharing = result.screenSharing;
        if (checkScreen && !supportScreenSharing) alert('browser is not support screenSharing');
        previewVideo = this.$refs.preview;
        start();
    } else {
        alert('不支持H264，请前往混流转码测试');
}

return true;
}

async function start() {
    initSDK();
    zg.setLogConfig({
        logLevel: 'debug',
        remoteLogLevel: 'info',
        logURL: '',
    });

    zg.setDebugVerbose(false);
    zg.setSoundLevelDelegate(true, 1000);

    this.$refs.createRoom.click(async () => {
        let loginSuc = false;
        try {
            loginSuc = await enterRoom();
            loginSuc && (await publish());
        } catch (error) {
            console.error(error);
        }
    });

    this.$refs.openRoom.click(async () => {
        await enterRoom();
    });

    this.$refs.leaveRoom.click(function() {
        logout();
    });

    this.$refs.stopPlaySound.click(() => {
        zg.setSoundLevelDelegate(false);
    });

    this.$refs.resumePlaySound.click(() => {
        zg.setSoundLevelDelegate(false);
        zg.setSoundLevelDelegate(true);
    });
}

async function enumDevices() {
    const audioInputList = [],
        videoInputList = [];
    const deviceInfo = await zg.enumDevices();

    deviceInfo &&
    deviceInfo.microphones.map((item, index) => {
        if (!item.deviceName) {
            item.deviceName = 'microphone' + index;
        }
        audioInputList.push(' <option value="' + item.deviceID + '">' + item.deviceName + '</option>');
        console.log('microphone: ' + item.deviceName);
        return item;
    });

    deviceInfo &&
    deviceInfo.cameras.map((item, index) => {
        if (!item.deviceName) {
            item.deviceName = 'camera' + index;
        }
        videoInputList.push(' <option value="' + item.deviceID + '">' + item.deviceName + '</option>');
        console.log('camera: ' + item.deviceName);
        return item;
    });

    audioInputList.push('<option value="0">禁止</option>');
    videoInputList.push('<option value="0">禁止</option>');

    this.$refs.audioList.html(audioInputList.join(''));
    this.$refs.videoList.html(videoInputList.join(''));
}

function initSDK() {
    enumDevices();

    zg.on('roomStateUpdate', (roomID, state, errorCode, extendedData) => {
        console.log('roomStateUpdate: ', roomID, state, errorCode, extendedData);
    });
    zg.on('roomUserUpdate', (roomID, updateType, userList) => {
        console.warn(
            `roomUserUpdate: room ${roomID}, user ${updateType === 'ADD' ? 'added' : 'left'} `,
            JSON.stringify(userList),
        );
    });
    zg.on('publisherStateUpdate', result => {
        console.log('publisherStateUpdate: ', result.streamID, result.state);
        if (result.state == 'PUBLISHING') {
            console.info(' publish  success ' + result.streamID);
        } else if (result.state == 'PUBLISH_REQUESTING') {
            console.info(' publish  retry');
        } else {
            if (result.errorCode == 0) {
                console.warn('publish stop ' + result.errorCode);
            } else {
                console.error('publish error ' + result.errorCode);
            }
            // const _msg = stateInfo.error.msg;
            // if (stateInfo.error.msg.indexOf ('server session closed, reason: ') > -1) {
            //         const code = stateInfo.error.msg.replace ('server session closed, reason: ', '');
            //         if (code === '21') {
            //                 _msg = '音频编解码不支持(opus)';
            //         } else if (code === '22') {
            //                 _msg = '视频编解码不支持(H264)'
            //         } else if (code === '20') {
            //                 _msg = 'sdp 解释错误';
            //         }
            // }
            // alert('推流失败,reason = ' + _msg);
        }
    });
    zg.on('playerStateUpdate', result => {
        console.log('playerStateUpdate', result.streamID, result.state);
        if (result.state == 'PLAYING') {
            console.info(' play  success ' + result.streamID);
        } else if (result.state == 'PLAY_REQUESTING') {
            console.info(' play  retry');
        } else {
            if (result.errorCode == 0) {
                console.warn('play stop ' + result.errorCode);
            } else {
                console.error('play error ' + result.errorCode);
            }

            // const _msg = stateInfo.error.msg;
            // if (stateInfo.error.msg.indexOf ('server session closed, reason: ') > -1) {
            //         const code = stateInfo.error.msg.replace ('server session closed, reason: ', '');
            //         if (code === '21') {
            //                 _msg = '音频编解码不支持(opus)';
            //         } else if (code === '22') {
            //                 _msg = '视频编解码不支持(H264)'
            //         } else if (code === '20') {
            //                 _msg = 'sdp 解释错误';
            //         }
            // }
            // alert('拉流失败,reason = ' + _msg);
        }
    });
    zg.on('roomStreamUpdate', async (roomID, updateType, streamList) => {
        console.log('roomStreamUpdate roomID ', roomID, streamList);
        if (updateType == 'ADD') {
            for (let i = 0; i < streamList.length; i++) {
                console.info(streamList[i].streamID + ' was added');
                useLocalStreamList.push(streamList[i]);
                let remoteStream;

                try {
                    remoteStream = await zg.startPlayingStream(streamList[i].streamID);
                } catch (error) {
                    console.error(error);
                    break;
                }

                this.$refs.remoteVideo.append($('<video  autoplay muted playsinline controls></video>'));
                const video = $('.remoteVideo video:last')[0];
                console.warn('video', video, remoteStream);
                video.srcObject = remoteStream;
                video.muted = false;
            }
        } else if (updateType == 'DELETE') {
            for (let k = 0; k < useLocalStreamList.length; k++) {
                for (let j = 0; j < streamList.length; j++) {
                    if (useLocalStreamList[k].streamID === streamList[j].streamID) {
                        try {
                            zg.stopPlayingStream(useLocalStreamList[k].streamID);
                        } catch (error) {
                            console.error(error);
                        }

                        console.info(useLocalStreamList[k].streamID + 'was devared');

                        useLocalStreamList.splice(k, 1);

                    }
                }
            }
        }
    });
}

async function login(roomId){
    // 获取token需要客户自己实现，token是对登录房间的唯一验证
    // Obtaining a token needs to be implemented by the customer. The token is the only verification for the login room.
    let token = '';

    token = await Promise.get('https://wsliveroom-alpha.zego.im:8282/token', {
        app_id: appID,
        id_name: userID,
    });

    return await zg.loginRoom(roomId, token, { userID, userName }, { userUpdate: true });
}

async function enterRoom() {
    await login("2");
    loginRoom = true;
    return true;
}

async function logout() {
    console.info('leave room  and close stream');
    // 停止推流
    // stop publishing
    if (isPreviewed) {
        zg.stopPublishingStream(publishStreamId);
        zg.destroyStream(localStream);
        isPreviewed = false;
        previewVideo.srcObject = null;
    }

    // 停止拉流
    // stop playing
    for (let i = 0; i < useLocalStreamList.length; i++) {
        useLocalStreamList[i].streamID && zg.stopPlayingStream(useLocalStreamList[i].streamID);
    }

    //退出登录
    //logout
    zg.logoutRoom("2");
    loginRoom = false;
}


export {
    zg,
    appID,
    publishStreamId,
    checkAnRun,
    supportScreenSharing,
    userID,
    useLocalStreamList,
    logout,
    enterRoom,
    previewVideo,
    isPreviewed,
    loginRoom,
    publishType,
};

