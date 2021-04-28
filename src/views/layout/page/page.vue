<template>
    <!-- 页面 -->
    <el-main>
        <div id="page">
            <vue-scroll :ops="ops">
                <div class="pageHeader">
                    <div class="title" @click="clickTitle">我的文档</div>
                    <div v-for="(item , index) in nodeTree" :key="item.id" @click="handleClick(item)">
                        <i class="el-icon-arrow-right icon"></i>
                        <span class="title"
                            :class="[nodeTree.length - 1 == index ? 'titlefocus':'']">{{item.label}}</span>
                    </div>
                </div>

                <div class="pageMain">
                    <div class="container" v-for="item in nodeList" :key="item.id" @click="handleClick(item)">
                        <div
                            style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;">
                            <i v-if="item.type" class="el-icon-folder foldIconColor"></i>
                            <i v-else class="el-icon-document doculmentIconColor"></i>
                            <div class="textItem">
                                <div style="font-size: 1.5em;">{{item.label}}</div>
                                <div style="color: #bababd;">4月25日 00:00</div>
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
                // //dialog
                // dialog: {
                //     dialogInput: '',// dialog输入文字
                //     dialogRename: false,//重命名dialog
                //     dialogDelete: false,//删除dialog
                // },
                //第二级点击的文件树菜单栏
                treeData: {},
                showAsideIcon: "el-icon-arrow-left", //侧边栏显示隐藏图标样式,
                ops: {
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
            };
        },
        computed: {
            nodeTree() {
                return this.$store.state.nodeTree;
            },
            nodeList() {
                return this.$store.state.nodeList;
            }
        },
        watch: {
        },
        methods: {
            clickTitle() {
                this.$emit("clickTitle");
            },
            handleClick(data) {
                // console.log("data", data)
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
                }

            },
            //点击二级菜单
            clickNode(item) { this.treeData = item },
            //侧边栏展示显示函数
            changeShowAsideIcon: function () {
                this.showAsideIcon =
                    this.showAsideIcon == "el-icon-arrow-left"
                        ? "el-icon-arrow-right"
                        : "el-icon-arrow-left";
                this.showAside();
            },
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