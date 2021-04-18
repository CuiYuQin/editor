<template>
  <div id="app">
    <transition name="el-fade-in">
      <router-view></router-view>
    </transition>
  </div>
</template>

<style>
body {
	margin: 0px;
	padding: 0px;
	font-family: Microsoft YaHei, Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB,  SimSun, sans-serif;
	font-size: 14px;
	-webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

#app {
	position: absolute;
	top: 0px;
	bottom: 0px;
	width: 100%;
  color: #2c3e50;
}
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
} */
</style>
