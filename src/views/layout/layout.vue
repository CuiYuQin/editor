<template>
  <div id="loyout">
    <!-- 布局 -->
    <el-container class="outContainer">
      <!-- 侧边栏 -->
      <transition name="slide-fade">
        <el-aside v-show="ifShowAside" style="width: 18%">文件 | 大纲</el-aside>
      </transition>

      <el-container>
        <!-- 主页面，即编辑器 -->
        <el-main>
          <div id="editor">
            <!-- 编辑器 -->
            <mavon-editor style="height: 100%"></mavon-editor>
            <!-- 侧边栏收缩按钮 -->
            <div class="footer">
              <i v-show="ifShowAside" class="el-icon-arrow-left" @click="showAside"> </i>
              <i v-show="!ifShowAside" class="el-icon-arrow-right" @click="showAside"> </i>
            </div>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
// import layoutAside from "./aside/aside";

// Local Registration
import { mavonEditor } from "mavon-editor";
import "mavon-editor/dist/css/index.css";

export default {
  name: "layout",
  components: {
    // 'v-layoutAside':layoutAside,
    mavonEditor,
    // or 'mavon-editor': mavonEditor
  },
  data: function () {
    return {
      ifShowAside: true,
    };
  },
  methods: {
    showAside: function () {
      this.ifShowAside = !this.ifShowAside;
    },
  },
};
</script>

<style scoped>
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

/* 外层布局容器 */
.outContainer {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: fixed;
}
/* 底部容器 */
.footer {
  background: rgba(170, 102, 102, 0.1);
  color: #060607;
  text-align: center;
  margin: 0px;
  padding: 0px;
  font-size: 24px;
  position: absolute;
  z-index: 9999;
  bottom: 6px;
}
/* 侧边栏容器 */
.el-aside {
  background-color: #d3dce6;
  text-align: center;
}
/* 主要区域容器 */
.el-main {
  background-color: #e9eef3;
  text-align: center;
  margin: 0px;
  padding: 0px;
}
/* 编辑器 */
#editor {
  margin: auto;
  height: 100%;
  width: 100%;
}
</style>
