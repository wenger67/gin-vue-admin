<template>
  <div id="data-view">
    <dv-full-screen-container>
      <top-header/>
      <span @click="sfByElement" class="time-clock">{{ nowDate}} - {{ nowTime}}</span>
      <div class="main-content">
        <dv-border-box-3 class="main-column">
          <ranking-board />
        </dv-border-box-3>
        <dv-border-box-3 class="main-column device-count">
          <center-cmp />
        </dv-border-box-3>
        <dv-border-box-3 class="main-column">
          <center-cmp />
        </dv-border-box-3>
      </div>
      <div class="bottom-nav">
        <dv-decoration-7 class="bottom-nav-item">首页</dv-decoration-7>
        <dv-decoration-7 class="bottom-nav-item">今日</dv-decoration-7>
        <dv-decoration-7 class="bottom-nav-item">月度</dv-decoration-7>
        <dv-decoration-7 class="bottom-nav-item">排行</dv-decoration-7>
      </div>
    </dv-full-screen-container>
  </div>
</template>

<script>
import topHeader from './component/topHeader'
import rankingBoard from './component/rankingBoard'
import centerCmp from './component/CenterCmp'

import sf from 'screenfull'

export default {
  name: 'DataView',
  components: {
    topHeader,
    rankingBoard,
    centerCmp
  },
  data () {
    return {
      nowDate:null,    //存放年月日变量
      nowTime:null,   //存放时分秒变量
      timer: "",           //定义一个定时器的变量
    }
  },
  filters: {
    formatDate:function(i) {
      return (i<10)?("0"+i) : i;
    },
  },
  methods: {
    formatTime(i) {
      return (i<10)?("0"+i) : i;
    },
    sfByElement() {
      const element = document.getElementById('data-view');
      if (element && sf.isEnabled) {
          sf.request(element);
      }
    },
    getTime(){
      const date = new Date();
      const year = date.getFullYear();
      const month = this.formatTime(date.getMonth() + 1);
      const day = this.formatTime(date.getDate());
      let hour= this.formatTime(date.getHours());
      const minute = this.formatTime(date.getMinutes());
      const second = this.formatTime(date.getSeconds());
      let str = "";
      if(hour>12) {
        hour -= 12;
        str = " PM";
      }else{
        str = " AM";                        
      }
      this.nowDate = year + "年" + month + "月" + day+"日";
      this.nowTime = hour + ":" + minute + ":" + second + str;
    },
  },
  created() {
    this.timer = setInterval(this.getTime, 1000);    
  },
  destroyed() {
    if (this.timer) clearInterval(this.timer)
  },
}
</script>

<style lang="less">
html, body {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
}

#data-view {
  width: 100%;
  height: 100%;
  background-color: #030409;
  color: #fff;

  #dv-full-screen-container {
    background-image: url('./component/img/bg.png');
    background-size: 100% 100%;
    box-shadow: 0 0 3px blue;
    display: flex;
    flex-direction: column;
  }

  .time-clock {
    position: absolute;
    font-size: 20px;
    font-weight: bold;
    left: 65%;
    margin-top: 64px;
  }

  .main-content {
    flex: 1;
    display: flex;
    height: 80%;
    flex-direction: row;
    margin-bottom: 48px;

    .border-box-content {
      padding: 20px;
      box-sizing: border-box;
      display: flex;
    }
  }

  .main-column {
    width: 33%;
  }

  .device-count {
    height: 50%;
  }

  .block-left-right-content {
    flex: 1;
    display: flex;
    margin-top: 20px;
  }

  .block-top-bottom-content {
    flex: 1;
    display: flex;
    flex-direction: column;
    box-sizing: border-box;
    padding-left: 20px;
  }

  .block-top-content {
    height: 55%;
    display: flex;
    flex-grow: 0;
    box-sizing: border-box;
    padding-bottom: 20px;
  }

  .bottom-nav {
    display: flex;
    flex-direction: row;
    height: 10%;
    width: 100%;
    box-sizing: border-box;
    justify-content: center;
    align-items: center;

    .bottom-nav-item {
      width: 150px;
      height: 48px;
    }
  }
}
</style>
