<template>
  <!-- 布局 -->
  <el-container class="outContainer">
    <!-- 侧边栏 -->
    <transition name="slide-fade">
      <v-aside v-show="ifShowAside"></v-aside>
    </transition>
    <!-- 主区域容器 -->
    <v-main :showAside="showAside"></v-main>
  </el-container>
</template>

<script>
import main from "@/views/layout/main/main.vue";
import aside from "@/views/layout/aside/aside.vue";

export default {
  name: "layout",
  components: {
    "v-aside": aside,
    "v-main": main,
  },
  data: function () {
    return {
      //侧边栏是否展示
      ifShowAside: true,
    };
  },
  methods: {
    //侧边栏展示
    showAside: function () {
      this.ifShowAside = !this.ifShowAside;
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
  position: fixed;
  background-color: #2c2d32;
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