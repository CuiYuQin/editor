<template>
    <!-- 页面 -->
    <el-main>
        <div id="page">
            <vue-scroll :ops="$store.state.vueScrolloOps">
                <div class="pageHeader">
                    <div class="title" @click="clickTitle(whichPage)">{{whichPage}}</div>
                    <div v-for="(item , index) in nodeTree" :key="item.id" @click="handleClick(item)">
                        <i class="el-icon-arrow-right icon"></i>
                        <span class="title"
                            :class="[nodeTree.length - 1 == index ? 'titlefocus':'']">{{item.label}}</span>
                    </div>
                </div>

                <div class="pageMain">
                    <!-- 我的文档 -->
                    <div v-show="whichPage == '我的文档'" class="container" v-for="item in nodeList" :key="item.id"
                        @click="handleClick(item)">
                        <div
                            style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;">
                            <i v-if="item.type" class="el-icon-folder foldIconColor"></i>
                            <i v-else class="el-icon-document doculmentIconColor"></i>
                            <div class="textItem">
                                <div style="font-size: 1.5em;">{{item.label}}</div>
                                <div style="color: #bababd;">{{item.time}}</div>
                            </div>
                        </div>

                        <el-dropdown trigger="click" placement='bottom-start' @command="handleCommand">
                            <i class="el-icon-more moreIcon" @click.stop="clickNode(item)"></i>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item icon="el-icon-edit" command="重命名">重命名</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-star-on" command="添加到快速访问">添加到快速访问</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-delete" command="删除">删除</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </div>

                    <!-- 最近编辑 -->
                    <div v-show="whichPage == '最近编辑'" class="container" v-for="item in editorTime" :key="item.id + '1'"
                        @click="handleClick(item)">
                        <div
                            style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;">
                            <i  class="el-icon-document doculmentIconColor"></i>
                            <div class="textItem">
                                <div style="font-size: 1.5em;">{{item.label}}</div>
                                <div style="color: #bababd;">{{item.time}}</div>
                            </div>
                        </div>

                        <el-dropdown trigger="click" placement='bottom-start' @command="handleCommand">
                            <i class="el-icon-more moreIcon" @click.stop="clickNode(item)"></i>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item icon="el-icon-edit" command="重命名">重命名</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-star-on" command="添加到快速访问">添加到快速访问</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-delete" command="删除">删除</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </div>

                    <!-- 回收站 -->
                    <div v-show="whichPage == '回收站'" class="container" v-for="(item, index) in nodeList"
                        :key="item.id + '2'" @click="handleDeleteClick(item)">
                        <div
                            style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;">
                            <i v-if="item.type" class="el-icon-folder foldIconColor"></i>
                            <i v-else class="el-icon-document doculmentIconColor"></i>
                            <div class="textItem">
                                <div style="font-size: 1.5em;">{{item.label}}</div>
                                <div style="color: #bababd;">{{item.time}}</div>
                            </div>
                        </div>

                        <el-dropdown trigger="click" placement='bottom-start' @command="handleCommand">
                            <i class="el-icon-more moreIcon" @click.stop="clickNode(item, index)"></i>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item icon="el-icon-magic-stick" command="恢复">恢复</el-dropdown-item>
                                <el-dropdown-item icon="el-icon-delete" command="彻底删除">彻底删除</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </div>
                </div>
            </vue-scroll>

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
    import vuescroll from "vuescroll"; //  引入vuescroll
    import "vuescroll/dist/vuescroll.css"; //  引入vuescroll样式
    export default {
        name: "pageContainer",
        components: {
            "vue-scroll": vuescroll,
        },
        props: {
            // 父组件传过来的侧边栏显示隐藏函数
            showAside: {
                type: Function,
            },
        },
        data: function () {
            return {
                //第二级点击的文件树菜单栏
                treeData: {},
                index: -1,
                showAsideIcon: "el-icon-arrow-left", //侧边栏显示隐藏图标样式,
            };
        },
        computed: {
            nodeTree() {
                return this.$store.state.nodeTree;
            },
            nodeList() {
                return this.$store.state.nodeList;
            },
            whichPage() {
                return this.$store.state.whichPage;
            },
            editorTime() {
                let data = this.$store.state.data;//获取根节点
                let queue = [];                   //队列，用于层级遍历时，缓存还需要层级遍历的节点
                let res = [];                     //返回的数组，就是文档节点数组
                //队列尾部拼接根节点，目前有一个节点
                queue = queue.concat(data);       
                //当队列有节点，就要继续层级遍历
                while (queue.length > 0) {
                    let node = queue[0]; //获取队列的数组头部节点
                    //移除队列的数组头部节点
                    queue.shift();    
                    //判断节点的类型 
                    if (!node.type) {  
                        //是文档就尾部加入返回数组  
                        res.push(node);
                    //是文件夹，就尾部加入孩子节点
                    } else {
                        queue = queue.concat(node.children)
                    }
                }
                //对文档节点按照编辑时间戳从大到小排序
                if (res.length > 1) {
                    res.sort(function (a, b) {
                        return b.editorTime - a.editorTime;
                    })
                }
                return res;
            }
        },
        watch: {
        },
        methods: {
            clickTitle(value) {
                this.$emit("clickTitle", value);
            },
            handleClick(data) {
                console.log("data", data)
                this.$emit("handleClick", data);
            },
            handleCommand(command) {
                console.log("click :", command);
                switch (command) {
                    case '重命名':
                        this.$emit("handleCommand", '重命名', this.treeData)
                        break;
                    case '添加到快速访问':
                        // this.$emit("handleCommand", '添加到快速访问')
                        break;
                    case '删除':
                        this.$emit("handleCommand", '删除', this.treeData)
                        break;
                    case '恢复':
                        this.$store.commit("spliceNodeList", this.index);
                        this.$emit("handleCommand", '恢复', this.treeData)
                        break;
                    case '彻底删除':
                        this.$store.commit("spliceNodeList", this.index);
                        this.$message({
                            showClose: true,
                            message: '已彻底删除 😊',
                            center: true,
                            duration: 1500,
                            type: 'success'
                        });
                        break;
                }

            },
            //点击二级菜单
            clickNode(item, index) {
                this.treeData = item;
                this.index = index;
            },
            //侧边栏展示显示函数
            changeShowAsideIcon: function () {
                this.showAsideIcon =
                    this.showAsideIcon == "el-icon-arrow-left"
                        ? "el-icon-arrow-right"
                        : "el-icon-arrow-left";
                this.showAside();
            },
            handleDeleteClick(item) {
                this.$message({
                    showClose: true,
                    message: '需要恢复 ' + item.label + ' 才能查看 😒',
                    center: true,
                    duration: 2500,
                    type: 'warning'
                });
            }
        },
    };
</script>

<style scoped>
    .pageHeader {
        position: sticky;
        background-color: #1D1D1F;
        color: #9D9DA6;
        cursor: pointer;
        font-size: 1.3em;
        top: 0px;
        padding: 40px 115px 40px 114px;
        margin-top: 60px;
        z-index: 1;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
    }

    .icon {
        margin: 0px 10px;
    }

    .title {
        padding: 1px 3px;
    }

    .titlefocus {
        color: #F4F4F5;
    }

    .title:hover {
        background-color: #353333;
        border-radius: .3em;
        color: #F4F4F5;
    }

    .pageMain {
        height: 100%;
        padding: 0px 115px 50px 115px;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
    }

    .container {
        position: relative;
        width: 390px;
        height: 80px;
        color: #F4F4F5;
        background-color: #1D1D1F;
        border-radius: .5em;
        border: 1px solid #66666D;
        margin: 12px 30px 12px 0px;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
        cursor: pointer;
    }

    .container:hover {
        background-color: #353333;
    }

    .textItem {
        height: 55px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;

    }

    .moreIcon {
        position: relative;
        right: 20px;
        margin: 2px;
        padding: 4px;
        border-radius: .3em;
        color: #E7E7E8 !important;
    }

    .moreIcon:hover {
        background-color: #49494E;
        color: #E7E7E8 !important;
    }


    /deep/ .foldIconColor {
        color: #774747;
        font-size: 2.5em;
        margin: 0px 15px;
    }

    /deep/ .doculmentIconColor {
        color: #D4D4D7;
        font-size: 2.5em;
        margin: 0px 15px;
    }



    .el-main {
        background-color: #1D1D1F;
        margin: 0px;
        padding: 0px;
        display: block;
        flex: 1;
        box-sizing: border-box;
    }

    /* 页面 */
    #page {
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
</style>