<template>
  <div>
    <el-aside style="transition:left .3s ease-in-out;">文件 | 大纲</el-aside>
        <el-container>
      <!-- 侧边栏 -->
      <!-- <layoutAside v-show="show"></layoutAside> -->
      <transition name="slide-fade">
        <el-aside v-show="show">文件 | 大纲</el-aside>
      </transition>
      <el-container>
        <!-- 主页 -->
        <el-main>
            <div id="editor">
              <mavon-editor></mavon-editor>
            </div>
        </el-main>
        <el-footer>
          <i class="el-icon-d-arrow-right" @click="showAside"> </i>Footer
        </el-footer>
      </el-container>
    </el-container>


      <div id="loyout">
    <!-- 布局 -->
    <el-container class="outContainer">
      <!-- 侧边栏 -->
      <el-aside>Aside</el-aside>
      <el-container>
        <!-- 主页面，即编辑器 -->
        <el-main>
          <div>
            <mavon-editor></mavon-editor>
          </div>
        </el-main>
        <!-- 底部 -->
        <el-footer height="30px">Footer</el-footer>
      </el-container>
    </el-container>
  </div>
  </div>
</template>

<script>
export default {
  name: "side",
};
</script>

<style>
.el-aside {
  background-color: #476b96;
  text-align: center;
  height: 100%;
}
</style>>


