<template>
  <div class="FM">
    <div class="el-collapse-item">
      <div class="el-collapse-item-header" :class="[collapse.activeName == '我的文档' ? 'isActive':'']"
        @click="clickCollapseItem('我的文档')">
        <div><i class="el-icon-s-management firstIcon"></i>我的文档</div><i :class="collapse.collapseTreeIcon"
          style="margin-right: 13px;" @click.stop="handleCollapseTree"></i>
      </div>
      <div class="el-collapse-item-wrap" :style="{display:!collapse.collapseTree ? '':'none'}">
        <vue-scroll :ops="$store.state.vueScrolloOps">
          <el-tree ref="tree" :props="defaultProps" :data="data" node-key="id" @node-drag-start="handleDragStart"
            :default-expanded-keys="expandedkeys" @node-drag-enter="handleDragEnter" @node-drag-leave="handleDragLeave"
            @node-drag-over="handleDragOver" @node-drag-end="handleDragEnd" @node-drop="handleDrop" draggable
            :allow-drag="allowDrag" :allow-drop="allowDrop" :indent="28" :expand-on-click-node='false'
            @node-click="handleNodeClick" @node-expand="handleNodeExpand">
            <!-- 自定义节点内容 -->
            <div class="custom-tree-node" slot-scope="{ node, data }">
              <div>
                <i v-if="data.type" class="el-icon-folder secondIcon foldIconColor"></i>
                <i v-else class="el-icon-document fourthIcon doculmentIconColor"></i>
                {{ node.label }}
              </div>
              <div>
                <!-- 新增下拉框 -->
                <el-dropdown v-if="data.type" trigger="click" placement='bottom-start' @command="handleCommand">
                  <i class="el-icon-plus thirdIcon" @click.stop="clickNode(data.id)"></i>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item icon="el-icon-folder foldIconColor " command="新建文件夹">新建文件夹</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-document " command="新建文档">新建文档</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
                <!-- 新增下拉框 -->

                <!-- 菜单下拉框 -->
                <el-dropdown trigger="click" placement='bottom-start' @command="handleCommand">
                  <i class="el-icon-more thirdIcon" @click.stop="clickNode(data.id)"></i>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item icon="el-icon-edit" command="重命名">重命名</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-star-on" command="添加到快速访问">添加到快速访问</el-dropdown-item>
                    <el-dropdown-item icon="el-icon-delete" command="删除">删除</el-dropdown-item>
                  </el-dropdown-menu>
                </el-dropdown>
                <!-- 菜单下拉框 -->
              </div>
            </div>
            <!-- 自定义节点内容 -->
          </el-tree>
        </vue-scroll>
      </div>
    </div>
    <div class="el-collapse-item" @click="clickCollapseItem('最近编辑')">
      <div class="el-collapse-item-header" :class="[collapse.activeName == '最近编辑' ? 'isActive':'']">
        <div><i class="el-icon-time firstIcon"></i>最近编辑</div>
      </div>
    </div>
    <div class="el-collapse-item" @click="clickCollapseItem('回收站')">
      <div class="el-collapse-item-header" :class="[collapse.activeName == '回收站' ? 'isActive':'']">
        <div><i class="el-icon-delete-solid firstIcon"></i>回收站</div>
      </div>
    </div>
    <!-- 新建文件夹对话框 -->
    <el-dialog title="新建文件夹" :visible.sync="dialog.dialogAddFold" width="25%" top='30vh'>
      <el-input placeholder="请输入名称" :autofocus="true" v-model="dialog.dialogInput" clearable>
      </el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialog.dialogAddFold = false">取 消</el-button>
        <el-button type="primary" :disabled="dialog.dialogInput == '' ? true:false" @click="addFlod">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 新建文件夹对话框 -->
    <!-- 重名名对话框 -->
    <el-dialog :title="treeData.type ? '重命名文件夹' : '重命名文档'" :visible.sync="dialog.dialogRename" width="25%" top='30vh'>
      <el-input placeholder="请输入名称" :autofocus="true" v-model="dialog.dialogInput" clearable>
      </el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialog.dialogRename = false">取 消</el-button>
        <el-button type="primary" :disabled="dialog.dialogInput == '' ? true:false" @click="rename">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 重名名对话框 -->
    <!-- 删除对话框 -->
    <el-dialog :title="treeData.type ? '删除文件夹' : '删除文档'" :visible.sync="dialog.dialogDelete" width="25%" top='30vh'>
      <div v-if="treeData.type" style="font-size: 1.03em;">确认删除此文件夹吗？删除后，所有人将无法继续访问此文件夹。</div>
      <div style="font-size: 1.03em;" v-else>确认删除此文档吗？删除后，所有人将无法继续访问此文档。</div>
      <div style="color:#92929C;margin-top: 15px;">删除后可通过回收站进行恢复<br>回收站文件将为您永久保留</div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialog.dialogDelete = false">取 消</el-button>
        <el-button type="primary" @click="deleteDialog">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 删除对话框 -->
  </div>

</template>

<script>
  import vuescroll from "vuescroll"; //  引入vuescroll
  import "vuescroll/dist/vuescroll.css"; //  引入vuescroll样式

  import axios from 'axios';

  export default {
    name: "FM",
    components: { "vue-scroll": vuescroll, },
    props: {},
    data: function () {
      return {
        collapse: {
          activeName: "我的文档",
          collapseTree: false,
          collapseTreeIcon: "el-icon-caret-top"
        },
        //dialog
        dialog: {
          dialogInput: '',// dialog输入文字
          dialogAddFold: false,//新建文件夹dialog
          dialogRename: false,//重命名dialog
          dialogDelete: false,//删除dialog
        },

        //第二级点击的文件树菜单栏
        treeData: {},
        //新增树节点id
        id: 1000,
        //默认展开的节点
        expandedkeys: [],
        //文件树数据结构
        data: [],
        object: {
          // id:"",
          // userName:this.$store.state.userName,
          // tree:this.data,
          // deleteTree:[],
          // createTime:""
        },
        timeout: null,
        // data: [{
        //   id: 1,
        //   label: '我的文档',
        //   type: true,
        //   url: `# markdown`,
        //   children: []
        // }
        // ],
        //
        defaultProps: {
          children: 'children',
          label: 'label'
        },
      };
    },
    computed: {
    },
    watch: {
      object: {
        handler() {
          clearTimeout(this.timeout);
          this.timeout = setTimeout(() => {
            //保存数据
            axios.post('/api/updatetree', this.object)
              .then(response => {
                console.log("/api/updatetree:", response.data.status);
              }).catch(err => {
                console.log(err);
              })
          }, 2000);
        },
        // immediate: true,
        deep: true
      }
    },
    created: function () {
      //请求数据
      axios.post('/api/gettree', {
        userName: this.$store.state.userName
      }).then(response => {
        this.object = response.data.object;
        this.object.tree = JSON.parse(response.data.object.tree);
        this.object.deleteTree = JSON.parse(response.data.object.deleteTree);
        this.data = this.object.tree;
        this.$store.commit("setNodeList", this.data);
        this.$store.commit("setIsShowMain", false);
        this.$store.commit("setData", this.data);
        console.log("/api/gettree:", response.data.status);
        // if()
      }).catch(err => {
        console.log(err);
      })
      // .finally(() => console.log("load......"))
    },

    methods: {
      // 一级菜单
      clickCollapseItem(value) {
        switch (value) {
          case "我的文档":
            this.collapse.activeName = '我的文档';
            this.$store.commit("setNodeList", this.data);
            this.$store.commit("setNodeTree", []);
            this.$store.commit("setIsShowMain", false);
            this.$store.commit("setWhichPage", "我的文档");
            break;
          case "最近编辑":
            this.collapse.activeName = '最近编辑';
            this.$store.commit("setNodeList", this.object.deleteTree);
            this.$store.commit("setNodeTree", []);
            this.$store.commit("setIsShowMain", false);
            this.$store.commit("setWhichPage", "最近编辑");
            this.$refs.tree.setCurrentKey();
            break;
          case "回收站":
            this.collapse.activeName = '回收站';
            this.$store.commit("setNodeList", this.object.deleteTree);
            this.$store.commit("setNodeTree", []);
            this.$store.commit("setIsShowMain", false);
            this.$store.commit("setWhichPage", "回收站");
            this.$refs.tree.setCurrentKey();
            break;
        }
      },
      handleCollapseTree() {
        if (!this.collapse.collapseTree) {
          this.collapse.collapseTreeIcon = "el-icon-caret-bottom";
          this.collapse.collapseTree = true;
          return
        }
        this.collapse.collapseTreeIcon = "el-icon-caret-top";
        this.collapse.collapseTree = false;
      },
      // 一级菜单


      //树形控件
      handleDragStart() { },
      handleDragEnter() { },
      handleDragLeave() { },
      handleDragOver() { },
      handleDragEnd() { },
      handleDrop(node) {
        this.$message({
          showClose: true,
          message: '移动成功😊',
          center: true,
          duration: 1000,
          type: 'success'
        });
        this.handleNodeClick(node.data);

      },
      allowDrop(draggingNode, dropNode, type) {
        if (dropNode.data.type) {
          return true;
        } else {
          if (type == 'inner') {
            return false;
          }
          return true;
        }
      },
      allowDrag() { return true; },

      //点击我是文档标题
      clickTitle(value) {
        this.clickCollapseItem(value);
      },

      handleNodeClick(data, clickNode) {
        let node = this.$refs.tree.getNode(data);
        //要展开的节点，也就是本节点
        this.expandedkeys = [data.id];
        //设置节点选中
        this.$refs.tree.setCurrentKey(data.id);
        this.treeData = data;
        if(!this.treeData.type) {
          this.setEditorTime();
        }
        if (data.type) {
          //nodeList
          let nodeList = data.children;
          this.$store.commit("setNodeList", nodeList);
          //nodetree
          let nodeTree = new Array(node.level);
          let i = node.level - 1;
          for (; node.parent; i--) {
            nodeTree[i] = node.data;
            node = node.parent;
          }
          this.$store.commit("setNodeTree", nodeTree);
          this.$store.commit("setIsShowMain", false);
          if (clickNode) {
            this.collapse.activeName = '我的文档';
            this.$store.commit("setWhichPage", "我的文档");
            return
          }
          if (this.collapse.activeName == '最近编辑') {
            // this.clickCollapseItem("最近编辑");
            return
          }
          if (this.collapse.activeName == '回收站') {
            this.clickCollapseItem("回收站");
            return
          }
          return
        }
        this.$store.commit("setIsShowMain", true);
        this.$store.commit("setValue", data.url);
        if (clickNode) {
          this.collapse.activeName = '我的文档';
          this.$store.commit("setWhichPage", "我的文档");
          return
        }
        if (this.collapse.activeName == '最近编辑') {
          // this.clickCollapseItem("最近编辑");
          return
        }
        if (this.collapse.activeName == '回收站') {
          this.clickCollapseItem("回收站");
          return
        }

      },

      handleNodeExpand() { },

      //点击二级菜单
      clickNode(id) { this.treeData = this.$refs.tree.getNode(id).data; },
      //处理二级下拉框菜单
      handleCommand(command) {
        switch (command) {
          case '新建文件夹':
            this.dialog.dialogInput = '';
            this.dialog.dialogAddFold = true;
            break;
          case '新建文档':
            this.addfile();
            break;
          case '重命名':
            this.dialog.dialogInput = this.treeData.label == 'Untitled' ? '' : this.treeData.label;
            this.dialog.dialogRename = true;
            break;
          // case '添加到快速访问': ;
          // break;
          case '删除':
            this.dialog.dialogDelete = true;
            break;
          case '恢复':
            this.data.push(this.treeData);
            this.collapse.collapseTreeIcon = "el-icon-caret-top";
            this.collapse.collapseTree = false;
            break;
        }

      },
      //二级菜单
      //生成时间
      CurentTime() {
        var now = new Date();

        var year = now.getFullYear();//年
        var month = now.getMonth() + 1;//月
        var day = now.getDate();//日

        var hh = now.getHours();//时
        var mm = now.getMinutes();//分
        var ss = now.getSeconds();//秒

        var clock = year + "-";
        if (month < 10)
          clock += "0";
        clock += month + "-";
        if (day < 10)
          clock += "0";
        clock += day + " ";
        if (hh < 10)
          clock += "0";
        clock += hh + ":";
        if (mm < 10)
          clock += '0';
        clock += mm + ":";
        if (ss < 10)
          clock += '0';
        clock += ss;
        return (clock);
      },

      //新建文件夹
      addFlod() {
        //关闭dialog
        this.dialog.dialogAddFold = false;
        //new 一个孩子
        const newChild = {
          id: new Date().getTime(),
          time: this.CurentTime(),
          editorTime:new Date().getTime(),
          label: this.dialog.dialogInput,
          type: true,
          url: '# ' + this.dialog.dialogInput,
          children: []
        };
        //如果本节点数据对象没有children这个属性，就要加入vue数据监听，要不不会响应式
        if (!this.treeData.children) {
          this.$set(this.treeData, 'children', []);
        }
        //把孩子节点加入孩子属性
        if (this.treeData === this.data) {
          this.treeData.push(newChild);
        } else {
          this.treeData.children.push(newChild);
        }
        this.$nextTick(function () {
          this.handleNodeClick(newChild);
        })
      },
      //新建文档
      addfile() {
        // 设置数据
        let ebookId = "";
        axios.post('/api/insertebook', {
          title: "Untitled",
          author: this.$store.state.userName,
          content: "## 请仔细阅读帮助文档",
          contentHtml: ""
        }).then(response => {
          console.log("/api/insertebook:", response.data.status);
          ebookId = response.data.object.ebookId;
          //new 一个孩子
          const newChild = {
            id: new Date().getTime(),
            time: this.CurentTime(),
            editorTime:new Date().getTime(),
            label: 'Untitled',
            type: false,
            url: ebookId,
            children: []
          };
          //如果本节点数据对象没有children这个属性，就要加入vue数据监听，要不不会响应式
          if (!this.treeData.children) {
            this.$set(this.treeData, 'children', []);
          }
          //把孩子节点加入孩子属性
          if (this.treeData === this.data) {
            this.treeData.push(newChild);
          } else {
            this.treeData.children.push(newChild);
          }
          this.$nextTick(function () {
            this.handleNodeClick(newChild);
          })
        }).catch(err => {
          console.log(err);
        })
      },
      //重命名
      rename() {
        //关闭dialog
        this.treeData.label = this.dialog.dialogInput;
        this.treeData.time = this.CurentTime();
        this.treeData.editorTime = new Date().getTime();
        this.dialog.dialogRename = false;
        if (!this.treeData.type) {
          axios.post('/api/updateebooktitle', {
            title: this.treeData.label,
            ebookId: this.treeData.url
          })
            .then(response => {
              console.log("/api/updateebooktitle:", response.data.status);
            }).catch(err => {
              console.log(err);
            })
        }
      },
      //删除
      deleteDialog() {
        let currentData = this.$refs.tree.getCurrentNode();
        let parent = this.$refs.tree.getNode(this.treeData).parent;
        //有选且为选中本节点，就删除后选它的父节点
        if (currentData && currentData.id == this.treeData.id && parent && parent.id != 0) {
          this.$refs.tree.setCurrentKey(parent.key);
        }

        this.$refs.tree.remove(this.treeData);
        //把该节点加入删除delereTree的头部
        this.object.deleteTree.unshift(this.treeData);

        this.dialog.dialogDelete = false;

        this.$nextTick(function () {
          currentData = this.$refs.tree.getCurrentNode();
          //有选择
          if (currentData) {
            let node = this.$refs.tree.getNode(currentData);
            //已被删除
            if (!node) {
              //有父母
              if (parent && parent.id != 0) {
                this.handleNodeClick(parent.data);
                //无父母
              } else {
                this.clickCollapseItem("我的文档");
              }
              //无被删除
            } else {
              this.handleNodeClick(currentData)
            }
            //无选择
          } else {
            //有父母
            if (parent && parent.id != 0) {
              this.handleNodeClick(parent.data);
              //无父母
            } else {
              this.clickCollapseItem("我的文档");
            }
          }
        })

      },
      //dialog
      setEditorTime() {
        this.treeData.editorTime = new Date().getTime();
      }
    },
  };
</script>

<style scoped>
  .FM {
    position: relative;
    height: 100%;
    display: flex;
    flex-direction: column;
  }


  /* 一级菜单 */
  .el-collapse-item {
    position: relative;
    max-height: 84%;
    display: flex;
    flex-direction: column;
  }

  .el-collapse-item-header {
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: space-between;
    height: 48px;
    line-height: 48px;
    color: #C9C9CE;
    cursor: pointer;
    font-size: 1.150em;
    font-weight: 500;
    border-radius: .7em;
    margin: 5px 10px;
  }

  .el-collapse-item-header:hover {
    background-color: #5856D5;
    color: rgb(255, 255, 255);
  }

  .el-collapse-item-wrap {
    background-color: #23292f;
    overflow: hidden;
    box-sizing: border-box;
  }

  .isActive {
    background-color: #5856D5;
    color: rgb(255, 255, 255);
  }

  /* 一级菜单 */


  /* 树形控件 */
  /deep/ .el-tree {
    position: relative;
    cursor: default;
    background: #23292f;
    color: #d9d9dcb5;
    padding: 0px;
  }

  /deep/ .el-tree-node {
    white-space: nowrap;
    outline: 0;
  }

  /deep/ .el-tree-node__content {
    display: flex;
    align-items: center;
    height: 30px;
    cursor: pointer;
    border-radius: .4em;
    margin: 0px 5px 0px 11px;
  }

  /deep/ .el-tree-node__expand-icon {
    color: #C0C4CC;
    font-size: .9em;
  }


  /deep/ .el-tree-node__content:hover {
    background-color: #606067;
    padding: 0px;
    border-radius: .4em;
  }

  /deep/ .el-tree-node:focus>.el-tree-node__content {
    background-color: #424244;
  }

  /deep/ .el-tree-node.is-current>.el-tree-node__content {
    background-color: #424244;
  }

  /deep/ .el-tree__drop-indicator {
    position: absolute;
    left: 0;
    right: 0;
    height: 3px;
    background-color: #ba9cc2;
  }

  /* 树形控件 */


  /* 新增下拉框 */
  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  /deep/ .el-dropdown-menu__item {
    list-style: none;
    line-height: 36px;
    padding: 0 20px;
    margin: 0;
    font-size: 14px;
    color: #EAEAEB;
    cursor: pointer;
    outline: 0;
  }

  /deep/ .el-dropdown-menu__item:focus,
  .el-dropdown-menu__item:not(.is-disabled):hover {
    background-color: #606067;
    color: #EAEAEB;
  }

  /* 新增下拉框 */

  /* dialog */

  /deep/ .v-modal {
    opacity: .75;
  }

  /deep/ .el-dialog {
    border-radius: 5px;
    box-shadow: 0 1px 3px rgb(0 0 0 / 80%);
    background: #49494E;
  }

  /deep/ .el-dialog__header {
    padding: 30px 30px 10px;
  }

  /deep/ .el-dialog__title {
    font-size: 1.5em;
    color: #F4F4F5;
  }

  /deep/ .el-dialog__headerbtn {
    top: 30px;
    right: 30px;
    font-size: 1.25em;
  }

  /deep/ .el-dialog__headerbtn .el-dialog__close {
    color: #E9E9EB;
  }

  /deep/ .el-dialog__headerbtn:hover .el-dialog__close {
    color: #E9E9EB;
  }

  /deep/ .el-dialog__body {
    padding: 20px 30px;
    color: #F4F4F5;
    font-size: 14px;
  }

  /deep/ .el-input__inner {
    background-color: #49494E;
    border-radius: 5px;
    border: 1px solid #58585D;
    color: #F4F4F5;
  }

  /* /deep/ .el-input__inner:hover {
    border-color: #3b38bb;
  } */

  /deep/ .el-dialog__footer {
    padding: 10px 30px 20px;
  }

  /deep/ .el-button {
    background: #92929C;
    border: 1px solid #dcdfe600;
    color: #ECECEE;
    font-weight: 500;
    padding: 10px 33px;
    font-size: 1.10em;
    border-radius: 5px;
  }

  /deep/ .el-button:focus,
  .el-button:hover {
    color: #ECECEE;
    border-color: #acacb0;
    background-color: #acacb0;
  }

  /deep/ .el-button--primary {
    color: #F4F4F5;
    background-color: #5856D5;
    border-color: #5856D5;
  }


  /deep/ .el-button--primary:focus,
  .el-button--primary:hover {
    background: #817fec;
    border-color: #817fec;
    color: #F4F4F5;
  }

  /* dialog */

  /* 图标 */

  .firstIcon {
    margin: 0px 15px;
  }


  /deep/ .foldIconColor {
    color: #774747;
  }

  /deep/ .doculmentIconColor {
    color: #D4D4D7;
  }

  .secondIcon {
    margin: 2px 10px;
  }

  .fourthIcon {
    margin: 2px 10px;
  }

  .thirdIcon {
    /* display: inline-block; */
    position: relative;
    margin: 0px 2px;
    color: #D4D4D7;
    padding: 3px 4px;
    border-radius: .2em;
    z-index: -999;
  }

  .custom-tree-node:hover .thirdIcon {
    z-index: 0;
  }

  .thirdIcon:hover {
    background-color: #49494E;
  }

  /* 图标 */
</style>

<style>
  /* 新增下拉框 */
  .el-dropdown-menu {
    background-color: #49494E !important;
    border: 0px !important;
    border-radius: 4px !important;
    box-shadow: 0 2px 12px 0 rgb(0 0 0 / 50%) !important;
  }

  .popper__arrow {
    border-style: solid !important;
    border-bottom-color: #ffffff00 !important;
  }

  .el-popper[x-placement^=bottom] .popper__arrow::after {
    border-bottom-color: #fff0 !important;
  }

  /* 新增下拉框 */
</style>