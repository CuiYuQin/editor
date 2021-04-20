<template>
  <!-- 侧边栏 -->
  <div class="aside">
    <!-- 头像区域 -->
    <div class="avatarArea">
      <div class="avaterContent">
        <!-- 头像 -->
        <el-avatar
          :size="avatarSize"
          :src="avatarUrl"
          style="vertical-align: middle"
        ></el-avatar>
        <!-- 昵称区域 -->
        <div
          class="nickStyle"
          :class="{ nickClick: ifnickClick }"
          @click="clickNick"
        >
          <a>{{ nickName }}</a> <i class="el-icon-caret-right"></i>
        </div>
      </div>
    </div>
    <!-- 标签区域 -->
    <el-tabs v-model="activeName" @tab-click="handleClick" :stretch="true">
      <!-- 文件标签页 -->
      <el-tab-pane label="文件" name="first">
        <vue-scroll :ops="ops">
          文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================

          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================

          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
          <br />文件========================================
        </vue-scroll>
      </el-tab-pane>

      <!-- 大纲标签页 -->
      <el-tab-pane label="大纲" name="second"
        ><vue-scroll :ops="ops"> ooooo </vue-scroll></el-tab-pane
      >
    </el-tabs>
  </div>
</template>

<script>
import vuescroll from "vuescroll"; //  引入vuescroll
import "vuescroll/dist/vuescroll.css"; //  引入vuescroll样式

export default {
  name: "asideContainer",
  components: {
    "vue-scroll": vuescroll,
  },
  data: function () {
    return {
      //头像大小 number / large / medium / small
      avatarSize: "large",
      //头像url
      avatarUrl:
        "https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg",
      //昵称名
      nickName: "starSea",
      //是否点击昵称
      ifnickClick: false,
      //标签页选定为第二页
      activeName: "first",
      ops: {
        vuescroll: {
          mode: "native", //选择一个模式, native 或者 slide(pc&app)
          sizeStrategy: "percent", //如果父容器不是固定高度，请设置为 number , 否则保持默认的percent即可
          detectResize: true, //是否检测内容尺寸发生变化
        },
        scrollPanel: {
          initialScrollY: false, //只要组件mounted之后自动滚动的距离。 例如 100 or 10%
          initialScrollX: false,
          scrollingX: true, //是否启用 x 或者 y 方向上的滚动
          scrollingY: true,
          speed: 300, //多长时间内完成一次滚动。 数值越小滚动的速度越快
          easing: undefined, //滚动动画 参数通animation
          verticalNativeBarPos: "right", //原生滚动条的位置
        },
        rail: {
          //轨道
          background: "#c3c3c3", //轨道的背景色
          opacity: 0,
          size: "6px",
          specifyBorderRadius: false, //是否指定轨道的 borderRadius， 如果不那么将会自动设置
          gutterOfEnds: null,
          gutterOfSide: "0px", //轨道距 x 和 y 轴两端的距离
          keepShow: false, //是否即使 bar 不存在的情况下也保持显示
        },
        bar: {
          showDelay: 1000, //在鼠标离开容器后多长时间隐藏滚动条
          onlyShowBarOnScroll: false, //当页面滚动时显示
          keepShow: false, //是否一直显示
          background: "#c3c3c3",
          opacity: 1,
          hoverStyle: false,
          specifyBorderRadius: false,
          minSize: false,
          size: "6px",
          disable: false, //是否禁用滚动条
        }, // 在这里设置全局默认配置
        name: "vuescroll", // 在这里自定义组件名字，默认是vueScroll
      },
    };
  },
  methods: {
    //点击昵称
    clickNick: function () {
      this.ifnickClick = !this.ifnickClick;
    },
    //标签页切换
    handleClick(tab, event) {
      console.log(tab, event);
    },
  },
};
</script>

<style scoped>
.aside {
  position: relative;
  background-color: #23292f;
  width: 18%;
  height: 100%;
  padding: 0px;
  margin: 0px;
  border: 1px solid red;
}
/* 头像区域 */
.avatarArea {
  position: relative;
  width: 100%;
  height: 8%;
  margin: 0px;
  padding: 0px;
}
.avaterContent {
  position: relative;
  top: 10px;
  left: 32px;
}
/* 昵称区域 */
.nickStyle {
  /* position: absolute; */
  font-size: 1.3em;
  font-weight: bold;
  display: inline-block;
  margin: 6px 15px;
  padding: 1px 5px 1px 6px;
  color: rgb(236, 229, 229);
}
/* 昵称点击 */
.nickClick {
  background-color: #49494e;
  border-radius: 4px;
}

/* 可以设置不同的进入和离开动画 */
/* 设置持续时间和动画函数 */
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.3s ease;
  /* transition:left .3s ease-in-out; */
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active for below version 2.1.8 */ {
  transform: translateX(-20px);
  opacity: 0;
}
</style>


<style>
/* 更改标签样式 */
.el-tabs {
  position: relative !important;
  width: 100% !important;
  height: 87% !important;
  margin: 0px !important;
  padding: 0px !important;
}
.el-tabs__header {
  position: absolute !important;
  padding: 0px !important;
  margin: 0px !important;
  width: 100% !important;
  height: 7% !important;
  top: 0px !important;
}
.el-tabs__content {
  position: absolute !important;
  padding: 0px !important;
  margin: 0px !important;
  width: 100% !important;
  height: 93% !important;
  bottom: 0px !important;
  color: #d0d0d5 !important;
  font-size: 1.19em !important;
}
.el-tab-pane {
  position: absolute !important;
  width: 100% !important;
  height: 100% !important;
  padding: 0px !important;
  margin: 0px !important;
}

.el-tabs__nav-wrap::after {
  background-color: #23292f !important;
}
.el-tabs__item {
  padding: 0 20px !important;
  font-size: 1.23em !important;
  font-weight: 400 !important;
  color: #787a80 !important;
}
.el-tabs__item.is-active {
  color: #ffffff !important;
  font-weight: 600 !important;
}
.el-tabs__active-bar {
  height: 2px !important;
  background-color: #ffffff !important;
}
</style>