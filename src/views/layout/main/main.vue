<template>
  <!-- 主页面，即编辑器和预览区 -->
  <el-main>
    <div id="editor">
      <!-- 编辑器 -->
      <mavon-editor style="height: 100%" v-model="d_value" :toolbarsBackground="'#23292f'"
        :previewBackground="'#f8f6f1ed'" :subfield="true" :placeholder="'请开始你的创作之旅...'" :toolbars="toolbars"
        :editorBackground="'#363B40'" @changeNav="changeNav"></mavon-editor>
      <!-- 侧边栏收缩按钮 -->
      <div class="footer" style="color: #ffffff">
        <el-tooltip class="item" effect="light" content="隐藏/显示侧边栏" placement="top">
          <!-- 图标 -->
          <i :class="showAsideIcon" @click="changeShowAsideIcon"></i>
        </el-tooltip>
      </div>
    </div>
  </el-main>
</template>

<script>
  // 引入编辑器
  import { mavonEditor } from "mavon-editor";
  import "mavon-editor/dist/css/index.css";

  import axios from 'axios';

  export default {
    name: "mainContainer",
    components: {
      mavonEditor,
    },
    props: {
      // 父组件传过来的侧边栏显示隐藏函数
      showAside: {
        type: Function,
      },
    },
    data: function () {
      return {
        showAsideIcon: "el-icon-arrow-left", //侧边栏显示隐藏图标样式
        d_value: "", //编辑文字
        object: {},
        //编辑器工具栏
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
      };
    },
    computed: {
      value() {
        return this.$store.state.value;
      },
    },
    watch: {
      value: {
        handler(newValue) {
          //请求数据
          axios.post('/api/getebook', {
            ebookId: newValue
          }).then(response => {
            console.log("response", response);
            this.object = response.data.object;
            this.d_value = this.object.content;
          }).catch(err => {
            console.log(err);
          })
        },
        immediate: true,
        // deep: true
      },
      d_value: {
        handler(newValue) {
          // //保存数据
          this.object.content = newValue;
          axios.post('/api/updateebook', this.object)
            .then(response => {
              // this.data = JSON.parse(response.data.object.tree)
              console.log(response.data.status);
            }).catch(err => {
              console.log(err);
            })
        },
        // immediate: true,
        deep: true
      }
    },
    methods: {
      //侧边栏展示显示函数
      changeShowAsideIcon: function () {
        this.showAsideIcon =
          this.showAsideIcon == "el-icon-arrow-left"
            ? "el-icon-arrow-right"
            : "el-icon-arrow-left";
        this.showAside();
      },
      changeNav: function ($vm) {
        this.$emit('changeNav', $vm);
      },
    },
  };
</script>

<style scoped>
  .el-main {
    background-color: #262d34;
    margin: 0px;
    padding: 0px;
    display: block;
    flex: 1;
    box-sizing: border-box;
  }

  /* 编辑器 */
  #editor {
    margin: auto;
    height: 100%;
    width: 100%;
  }

  /* 侧边栏按钮 */
  .footer {
    background: rgba(203, 19, 65, 0.5);
    color: #ffffff;
    text-align: center;
    margin: 0px;
    padding: 0px;
    font-size: 24px;
    position: absolute;
    z-index: 5000;
    bottom: 6px;
    cursor: pointer;
  }
</style>