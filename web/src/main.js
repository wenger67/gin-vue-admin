import Vue from 'vue'
import App from './App.vue'
// 引入element
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// 全局配置elementui的dialog不能通过点击遮罩层关闭
ElementUI.Dialog.props.closeOnClickModal.default = false
Vue.use(ElementUI);
// 引入封装的router
import router from '@/router/index'

// amap
import VueAMap from 'vue-amap';
Vue.use(VueAMap);
VueAMap.initAMapApiLoader({
    key:'2ace90bc407f7511c1628692685b0c5c',
    plugin:['AMap.Autocomplete', 'AMap.PlaceSearch', 'AMap.Scale', 'AMap.OverView', 'AMap.ToolBar', 'AMap.MapType', 'AMap.PolyEditor', 'AMap.CircleEditor'],
    v: '1.4.4'
})

// canvas背景插件
import vueParticleLine from 'vue-particle-line'
import 'vue-particle-line/dist/vue-particle-line.css'
Vue.use(vueParticleLine)

// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'

// 富文本插件
import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
Vue.use(VueQuillEditor)

// markdown插件
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(mavonEditor)

import '@/permission'
import { store } from '@/store/index'
Vue.config.productionTip = false

// 路由守卫
import Bus from '@/utils/bus.js'
Vue.use(Bus)

import {auth} from '@/directive/auth'
// 按钮权限指令
auth(Vue)

// http request
import axios from 'axios'
Vue.prototype.$http = axios

// websocket
import VueNativeSocket from 'vue-native-websocket'
Vue.use(VueNativeSocket, 'ws://127.0.0.1:8888/api/ws/endpoint', {
    format: 'json',
    reconnection: true,
    reconnectionAttempts: 5,
    reconnectionDelay: 3000
})

new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')

//引入echarts
import echarts from 'echarts'
Vue.prototype.$echarts = echarts;
