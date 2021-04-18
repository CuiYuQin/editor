<template>
  <div id="loyout">
    <!-- 布局 -->
    <el-container class="outContainer">
      <!-- 侧边栏 -->
      <transition name="slide-fade">
        <el-aside v-show="ifShowAside" style="width: 18%">

          <el-divider style="top:30%"></el-divider>
          <el-tabs v-model="activeName" @tab-click="handleClick" style="text-align: center;  margin: 0px;
  padding: 0px;">
            <el-tab-pane label="用户管理" name="first">用户管理</el-tab-pane>
            <el-tab-pane label="配置管理" name="second">配置管理</el-tab-pane>
          </el-tabs>

        </el-aside>
      </transition>

      <el-container>
        <!-- 主页面，即编辑器 -->
        <el-main>
          <div id="editor">
            <!-- 编辑器 -->
            <mavon-editor
              style="height: 100%"
              v-model="value"
              :toolbarsBackground="'#23292f'"
              :previewBackground="'#f8f6f1ed'"
              :subfield="true"
              :placeholder="'请开始你的创作之旅...'"
              :toolbars="toolbars"
              :editorBackground="'#2c2d32'"
            ></mavon-editor>
            <!-- 侧边栏收缩按钮 -->
            <div class="footer" style="color:#ffffff">
              <i
                v-show="ifShowAside"
                class="el-icon-arrow-left"
                @click="showAside"
              >
              </i>
              <i
                v-show="!ifShowAside"
                class="el-icon-arrow-right"
                @click="showAside"
              >
              </i>
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
      //侧边栏===================
      ifShowAside: true,
      activeName: 'second',
      //========================
      //markdown================
      value: ``,
      toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        strikethrough: true, // 中划线
        mark: true, // 标记
        superscript: true, // 上角标
        subscript: true, // 下角标
        quote: true, // 引用
        ol: true, // 有序列表
        ul: true, // 无序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        code: true, // code
        table: true, // 表格
        fullscreen: false, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        htmlcode: true, // 展示html源码
        help: false, // 帮助
        /* 1.3.5 */
        undo: true, // 上一步
        redo: true, // 下一步
        trash: true, // 清空
        save: true, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true, // 导航目录
        /* 2.1.8 */
        alignleft: true, // 左对齐
        aligncenter: true, // 居中
        alignright: true, // 右对齐
        /* 2.2.1 */
        subfield: true, // 单双栏模式
        preview: false, // 预览
      },
      //========================
    };
  },
  methods: {
    //侧边栏=====================================================================
    showAside: function () {
      this.ifShowAside = !this.ifShowAside;
    },
    handleClick(tab, event) {
        console.log(tab, event);
    }
    //==========================================================================
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
  /* background-color: #d3dce6; */
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
