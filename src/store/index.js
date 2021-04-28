import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    value: "",//markdown编辑文本
    isShowMain: false,
    nodeTree: [],
    nodeList: [],

  },
  mutations: {
    setValue(state, value) {
      state.value = value;
    },
    setIsShowMain(state, value) {
      state.isShowMain = value;
    },
    setNodeTree(state, nodeTree) {
      state.nodeTree = nodeTree;
    },
    setNodeList(state, nodeList) {
      state.nodeList = nodeList;
    }
  },
  actions: {
  },
  modules: {
  }
})
