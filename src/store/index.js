import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    userName: "",//用户名
    token:"",
    value: "",//markdown编辑文本
    isShowMain: false,
    whichPage:"我的文档",
    nodeTree: [],
    nodeList: [],
    data:[],
    ggbHtml:"",

    vueScrolloOps: {
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
        gutterOfEnds: "0px",
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

  },
  mutations: {
    setToken(state, userName) {
      state.token = userName;
    },
    setUserName(state, userName) {
      state.userName = userName;
    },
    setValue(state, value) {
      state.value = value;
    },
    setWhichPage(state, whichPage) {
      state.whichPage = whichPage;
    },
    setIsShowMain(state, isShowMain) {
      state.isShowMain = isShowMain;
    },
    setNodeTree(state, nodeTree) {
      state.nodeTree = nodeTree;
    },
    setNodeList(state, nodeList) {
      state.nodeList = nodeList;
    },
    setData(state, data) {
      state.data = data;
    },
    spliceNodeList(state,index) {
      state.nodeList.splice(index,1)
    },
    setVuescrolloOps(state, vuescrolloOps) {
      state.vuescrolloOps = vuescrolloOps
    },
    setGgbHtml(state, value) {
      state.ggbHtml = value
    }
  },
  actions: {
  },
  modules: {
  }
})
