<template>
    <div>
        <div>
            <el-button type="primary" @click="requestPush()">请求推流</el-button>
            <el-button type="primary" @click="startVideoCall()">视频监控</el-button>
        </div>
        <div>
            <el-row :gutter="15">
                <el-col :span="24">
                    <el-tag v-for="item in dynamicTags"
                            :key="item.label"
                            :type="item.type"
                            style="margin: 8px 5px;"
                            effect="plain">
                        {{ item.label }}
                    </el-tag>
                </el-col>
                <el-col :span="12" style="margin-top: 20px">
                    <video autoplay :key="index" v-for="(vd,index) in streamList" :srcObject.prop="vd.stream"/>
                </el-col>
                <el-col :span="12" style="margin-top: 20px">
                    <el-amap vid="amapDemo" :center="mapCenter" :amap-manager="aMapManager" :zoom="12" :events="events" class="amap-demo">
                        <el-amap-circle
                                v-for="(circle, index) in circles"
                                :center="circle.center"
                                :radius="circle.radius"
                                :key="index"
                                :fill-opacity="circle.fillOpacity"
                                :events="circle.events">
                        </el-amap-circle>
                    </el-amap>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
    import {
        findLift,
    } from "@/api/lift";
    import { mapGetters } from 'vuex';

    import VueAMap from "vue-amap";
    import {formatOnline, formatDate, randomType} from "@/utils/stringFun";
    let aMapManager = new VueAMap.AMapManager();
    import { ZegoExpressEngine } from 'zego-express-engine-webrtc';

    export default {
        name: "LiftDetial",
        data() {
            return {
                aMapManager,
                events: {
                    init(o) {
                        // eslint-disable-next-line no-undef
                        let marker = new AMap.Marker({
                            position: [121.59996, 31.197646]
                        });
                        marker.setMap(o);
                    }
                },
                circles: [
                    {
                        center: [121.5273285, 31.21515044],
                        radius: 200,
                        fillOpacity: 0.5,
                        events: {
                        }
                    }
                ],
                searchOption: {
                    city: '上海',
                    citylimit: true
                },
                mapCenter: [121.59996, 31.197646],
                dynamicTags: [],
                curLift: {},
                appID: 3483953860,
                server: 'wss://webliveroom-test.zego.im/ws',
                ze: undefined,
                userID: 'sample' + new Date().getTime(),
                userName:'sampleUser' + new Date().getTime(),
                streamList:[]
            };
        },
        computed: {
          ...mapGetters("user", ["userInfo"])
        },
        methods: {
            async handleAddressChange() {
                let address = this.curLift.address
                // change map center
                this.searchOption.city = address.region.city
                let location = this.curLift.location.split(',')
                this.mapCenter = [location[0], location[1]]
                this.circles[0].center = this.mapCenter
                await this.aMapManager.getMap().setZoom(16)
            },
            async getDetail(id) {
                this.dynamicTags = []
                let res = await findLift({ID:id})
                if (res.code === 0) {
                    let lift = res.data.relift
                    this.curLift = lift
                    this.dynamicTags.push({label: '别名:' + lift.nickName, type:randomType()})
                    this.dynamicTags.push({label: '编号:' + lift.code, type:randomType()})
                    this.dynamicTags.push({label: '安装公司:' + lift.installer.fullName, type:randomType()})
                    this.dynamicTags.push({label: '维保公司:' + lift.maintainer.fullName, type:randomType()})
                    this.dynamicTags.push({label: '年检公司:' + lift.checker.fullName, type:randomType()})
                    this.dynamicTags.push({label: '使用单位:' + lift.owner.fullName, type:randomType()})
                    this.dynamicTags.push({label: '出厂时间:' + formatDate(lift.factoryTime), type:randomType()})
                    this.dynamicTags.push({label: '安装时间:' + formatDate(lift.installTime), type:randomType()})
                    this.dynamicTags.push({label: '安检时间:' + formatDate(lift.checkTime), type: randomType()})
                    this.dynamicTags.push({label: '电梯型号:' + lift.liftModel.brand + '-' + lift.liftModel.modal, type: randomType})
                    this.dynamicTags.push({label: '使用场景:' + lift.category.categoryName, type: randomType()})
                    this.dynamicTags.push({label: '地址:' + lift.address.addressName + lift.building + '栋' + lift.cell + '单元', type: randomType()})
                    this.dynamicTags.push({label: '总楼层:' + lift.floorCount, type: randomType()})
                    this.dynamicTags.push({label: '智能设备:' + lift.adDevice.type.categoryName, type: randomType()})
                    this.dynamicTags.push({label: '智能设备状态:' + formatOnline(lift.adDevice.online), type: randomType()})
                }
            },
            initSDK() {
                this.zg.on('roomStateUpdate', (roomID, state, errorCode, extendedData) => {
                    console.log('roomStateUpdate: ', roomID, state, errorCode, extendedData);
                });
                this.zg.on('roomUserUpdate', (roomID, updateType, userList) => {
                    console.warn(
                        `roomUserUpdate: room ${roomID}, user ${updateType === 'ADD' ? 'added' : 'left'} `,
                        JSON.stringify(userList),
                    );
                });
                this.zg.on('publisherStateUpdate', result => {
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
                    }
                });
                this.zg.on('playerStateUpdate', result => {
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
                this.zg.on('roomStreamUpdate', async (roomID, updateType, streamList) => {
                    console.log('roomStreamUpdate roomID ', roomID, streamList);
                    if (updateType == 'ADD') {
                        for (let i = 0; i < streamList.length; i++) {
                            console.info(streamList[i].streamID + ' was added');
                            // useLocalStreamList.push(streamList[i]);
                            let remoteStream;
                            try {
                                remoteStream = await this.zg.startPlayingStream(streamList[i].streamID);
                            } catch (error) {
                                console.error(error);
                                break;
                            }
                            this.streamList.push({stream: remoteStream})
                        }
                    } else if (updateType == 'DELETE') {
                        // for (let k = 0; k < useLocalStreamList.length; k++) {
                        //     for (let j = 0; j < streamList.length; j++) {
                        //         if (useLocalStreamList[k].streamID === streamList[j].streamID) {
                        //             try {
                        //                 zg.stopPlayingStream(useLocalStreamList[k].streamID);
                        //             } catch (error) {
                        //                 console.error(error);
                        //             }
                        //
                        //             console.info(useLocalStreamList[k].streamID + 'was devared');
                        //
                        //             useLocalStreamList.splice(k, 1);
                        //
                        //         }
                        //     }
                        // }
                    }
                });
            },
            async login(roomId) {
                // 获取token需要客户自己实现，token是对登录房间的唯一验证
                // Obtaining a token needs to be implemented by the customer. The token is the only verification for the login room.

                let userID = this.userID
                let userName = this.userName

                let token = '';
                token = await this.$http.get('https://wsliveroom-alpha.zego.im:8282/token',{
                    params: {
                        app_id: this.appID,
                        id_name: userID
                    }
                });

                console.log('login:' + JSON.stringify(token.data))
                return await this.zg.loginRoom(roomId, token.data, { userID, userName }, { userUpdate: true });
            },
            async startVideoCall() {
                await this.login("2");

            },
            requestPush() {
                this.$socket.sendObj({
                    target: ["LIFT_" + this.curLift.ID],
                    event: "push.stream",
                    data: this.userInfo.uuid + ""
                })
            }
        },
        async created() {
            let id = this.$route.params.id
            await this.getDetail(id)
            await this.handleAddressChange()
            this.zg = new ZegoExpressEngine(this.appID, this.server);
            this.initSDK();
            this.$options.sockets.onmessage = async (data) => {
                let json = JSON.parse(data.data)
                if (json.event === "push.stream.success") {
                    let roomId = (json.data).split(":")[1]
                    await this.login(roomId);
                }
            }
        },
        beforeDestroy() {
            delete this.$options.sockets.onmessage
            this.zg.stopPlayingStream("LIFT_" + this.curLift.ID)
            this.zg.logoutRoom("LIFT_" + this.curLift.ID)
        }
    };
</script>

<style>
    .amap-demo {
        height: 300px!important;
        padding: 10px 20px 20px 20px;
    }

    .search-box {
        position: relative;
    }
</style>