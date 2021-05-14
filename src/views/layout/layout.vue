<template>
  <!-- 布局 -->
  <el-container class="outContainer">
    <!-- 侧边栏 -->
    <transition name="slide-fade">
      <v-aside v-show="ifShowAside" ref="asideRef"></v-aside>
    </transition>
    <!-- 主区域容器 -->

    <v-main v-show="isShowMain" :showAside="showAside" @changeNav='changeNav' @setEditorTime="setEditorTime" ref="mainRef"></v-main>
    <v-page v-show="!isShowMain" :showAside="showAside" ref="pageRef" @clickTitle="clickTitle"
      @handleClick="handleClick" @handleCommand="handleCommand"></v-page>


  </el-container>
</template>

<script>
  import main from "@/views/layout/main/main.vue";
  import aside from "@/views/layout/aside/aside.vue";
  import page from "@/views/layout/page/page.vue";

  export default {
    name: "layout",
    components: {
      "v-aside": aside,
      "v-main": main,
      "v-page": page,
    },
    data: function () {
      return {
        //侧边栏是否展示
        ifShowAside: true,
      };
    },
    computed: {
      //主区域
      isShowMain() {
        return this.$store.state.isShowMain;
      },
    },
    created: function () {
      if (this.$store.state.token == '' || this.$store.state.token != this.$store.state.userName) {
        this.$router.push({ path: "/login" })
      }
    },
    methods: {
      //侧边栏展示
      showAside: function () {
        this.ifShowAside = !this.ifShowAside;
      },
      clickTitle(value) {
        this.$refs.asideRef.$refs.FMRef.clickTitle(value);
      },
      handleClick(data) {
        this.$refs.asideRef.$refs.FMRef.treeData = data;
        this.$refs.asideRef.$refs.FMRef.handleNodeClick(data);
      },
      handleCommand(type, data) {
        this.$refs.asideRef.$refs.FMRef.treeData = data;
        this.$refs.asideRef.$refs.FMRef.handleCommand(type);
      },
      //
      //设置最近编辑时间
      setEditorTime() {
        this.$refs.asideRef.$refs.FMRef.setEditorTime();
        console.log("editorTime:",this.$refs.asideRef.$refs.FMRef.treeData)
      },
      //解析渲染大纲
      changeNav: function ($vm) {

        // changeNav: function () {
        let navigationContent;
        navigationContent = this.$refs.asideRef.$refs.navigationContent;
        navigationContent.innerHTML = $vm.d_render;
        let nodes = navigationContent.children;
        if (nodes.length) {
          for (let i = 0; i < nodes.length; i++) {
            judageH(nodes[i], i, nodes);
          }
        }
        function judageH(node, i) {
          let reg = /^H[1-6]{1}$/;
          if (!reg.exec(node.tagName)) {
            node.style.display = "none";
          } else {
            node.onclick = function () {
              let vShowContent = $vm.$refs.vShowContent;
              let vNoteEdit = $vm.$refs.vNoteEdit;
              if ($vm.s_subfield) {
                // 双栏
                if ($vm.s_preview_switch) {
                  // 编辑预览
                  vNoteEdit.scrollTop =
                    (vShowContent.children[i].offsetTop *
                      vNoteEdit.scrollHeight) /
                    vShowContent.scrollHeight;
                } 
              } else {
                // 单栏
                if ($vm.s_preview_switch) {
                  // 预览
                  vShowContent.scrollTop = vShowContent.children[i].offsetTop;
                } else {
                  // todo 编辑
                }
              }
            };
          }
        }
      },

    },
  };
</script>

<style scoped>
  .outContainer {
    height: 100%;
    width: 100%;
    margin: 0px;
    padding: 0px;
    position: absolute;
    background-color: #1D1D1F;
    overflow: hidden;
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

  .slide-fade-enter,
  .slide-fade-leave-to

  /* .slide-fade-leave-active for below version 2.1.8 */
    {
    transform: translateX(-20px);
    opacity: 0;
  }
</style>