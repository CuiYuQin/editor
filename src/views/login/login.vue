<template>
  <div id="login">
    <div class="loginConbox">
      <div class="header">
        <!--<div class="logo">-->
        <!--<img src="../../assets/logo.png">-->
        <!--</div>-->
      </div>
      <div class="loginBox">
        <div class="loginCon">
          <div>
            <p class="title">基于Web的交互电子书内容生成系统</p>
            <p class="content">内容生成器本质上是一个在线编辑器，除了可以进行常规的文档编辑，还提供了基于markdown语法的写作编辑。
              这个系统将教师和学生编辑的内容多姿多彩地展现，可以大大激发创造热情，丰富平台教学资源内容。</p>
          </div>

          <el-card shadow="always" class="login-module" v-if="smdl">
            <div slot="header" class="clearfix formTitlt">
              <span>密码登录</span>
              <!-- <span class="titIconbox">
                <i class="iconfont xu-saomadenglu2 fa-lg iconcolor"></i>
                <i class="iconfont xu-saomadenglu01 el-icon--right fa-lg pointer" @click="smdl = !smdl"></i>
              </span> -->
            </div>
            <el-form :model="loginForm" status-icon label-width="100px" class="demo-ruleForm">
              <el-form-item>
                <el-input type="text" v-model="loginForm.userName" auto-complete="off" placeholder="请输入登录账号"></el-input>
              </el-form-item>
              <el-form-item>
                <el-input type="password" v-model="loginForm.PassWord" auto-complete="off" placeholder="请输入登录密码">
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button class="subBtn" type="primary" @click="submitForm(1)">登录</el-button>
              </el-form-item>
              <p class="smalltxt">
                <!-- <router-link class="a" to="#">忘记密码</router-link>
                <router-link class="a" to="#">忘记会员名</router-link> -->
                <span class="a" @click="smdl = !smdl">免费注册</span>
              </p>
            </el-form>
          </el-card>

          <el-card shadow="always" class="login-module" v-else>
            <div slot="header" class="clearfix formTitlt">
              <span>免费注册</span>
              <!-- <span class="titIconbox">
                <i class="iconfont xu-saomadenglu2 fa-lg iconcolor"></i>
                <i class="iconfont xu-saomadenglu01 el-icon--right fa-lg pointer" @click="smdl = !smdl"></i>
              </span> -->
            </div>
            <el-form :model="loginForm" status-icon label-width="100px" class="demo-ruleForm">
              <el-form-item>
                <el-input type="text" v-model="loginForm.userName" auto-complete="off" placeholder="请输入注册账号"></el-input>
              </el-form-item>
              <el-form-item>
                <el-input type="password" v-model="loginForm.PassWord" auto-complete="off" placeholder="请输入注册密码">
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button class="subBtn" type="primary" @click="submitForm(2)">注册</el-button>
              </el-form-item>
              <p class="smalltxt">
                <!-- <router-link class="a" to="#">忘记密码</router-link>
                <router-link class="a" to="#">忘记会员名</router-link> -->
                <span class="a" @click="smdl = !smdl">密码登录</span>
              </p>
            </el-form>
          </el-card>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import axios from 'axios';

  export default {
    name: "login",
    data() {
      return {
        smdl: true,
        loginForm: {
          // userName: "starSea",
          // PassWord: "e10adc3949ba59abbe56e057f20f883e"
          userName: "",
          PassWord: ""
        }
      }
    },
    methods: {
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
      submitForm(num) {
        let that = this
        if (this.loginForm.userName === "" || this.loginForm.PassWord === "") {
          this.$message({
            showClose: true,
            message: "账号或密码不能为空",
            type: "error"
          })
          return false
        } else {
          //登录
          if (num == 1) {
            axios.post('/api/login', {
              userName: that.loginForm.userName,
              PassWord: that.loginForm.PassWord
            }).then(res => {
              console.log("/api/login:", res.data.status);
              if (res.data.status == 200) {
                that.$store.commit("setToken", that.loginForm.userName)
                that.$store.commit("setUserName", that.loginForm.userName)
                that.$router.push({ path: "/editor" })
              } else if (res.data.status == 300) {
                that.$message({
                  showClose: true,
                  message: "密码错误，请重新登录！",
                  type: "error"
                })
                this.PassWord = "";
              } else {
                that.$message({
                  showClose: true,
                  message: "不存在这个账号，请先注册！",
                  type: "warn"
                })
              }
            }).catch((err) => {
              console.log(err)
            })
          }
          //注册
          if (num == 2) {
            axios.post('/api/inserttree', {
              userName: that.loginForm.userName,
              PassWord: that.loginForm.PassWord,
              types: 1,
              tree: [
                {
                  id: new Date().getTime(),
                  time: this.CurentTime(),
                  editorTime: new Date().getTime(),
                  label: "文件夹",
                  type: true,
                  url: '# 文件夹',
                  children: []
                }
              ],
              deleteTree: []
            }).then(res => {
              console.log("/api/inserttree:", res.data.status);
              if (res.data.status == 200) {
                that.$store.commit("setToken", that.loginForm.userName)
                that.$store.commit("setUserName", that.loginForm.userName)
                that.$router.push({ path: "/editor" })
              } else {
                that.$message({
                  showClose: true,
                  message: "已有本账号，请用密码登录！",
                  type: "warn"
                })
              }
            }).catch((err) => {
              console.log(err)
            })
          }
        }
      },
    },

  }
</script>
<style lang="scss">
  #login {
    width: 100%;
    height: 100%;
    background-image:url("https://api.dujin.org/bing/1920.php");
    position: fixed;

    .header {
      height: 60px;
      position: relative;

      /*border-bottom: 1px solid rgba(255, 255, 255, 0.3);*/
      .logo {
        margin-left: 30px;
        width: 500px;
        float: left;
        height: 40px;
        padding-top: 10px;

        img {
          height: 100%;
        }
      }
    }

    .loginBox {
      .iconcolor {
        color: #409EFF;
      }

      padding: 74px 0 118px;

      .loginCon {
        width: 900px;
        margin: auto;
        position: relative;
        height: 388px;

        .el-card__header {
          border-bottom: 0px;
        }

        .title {
          font-size: 1.8em;
          font-weight: 600;
          color: #eef0e4;
          width: 520px;
          float: left;
          margin-top: 0px;

          &:first-child {
            font-size: 2.0em;
            margin-top: 50px;
            margin-bottom: 30px;
          }
        }

        .content {
          font-size: 1.2em;
          font-weight: 400;
          color: #e8ebee;
          width: 430px;
          float: left;
          margin-top: 10px;
        }

        .login-module {
          width: 380px;
          height: 325px;
          margin-top: 60px;
          border: none;
          position: absolute;
          right: 0;

          .formTitlt {
            font-size: 18px;
            font-weight: 400;

            .titIconbox {
              float: right;

              .pointer {
                cursor: pointer;
              }
            }
          }

          .smalltxt {
            text-align: right;

            .a {
              text-decoration: none;
              color: #999999;
              font-size: 12px;
              margin-left: 8px;
              cursor: pointer;
            }
          }
        }


        .el-form-item__content {
          margin-left: 0px !important;

          .subBtn {
            width: 100%;
          }
        }
      }

      .el-input__inner,
      .el-button,
      .el-card,
      .el-message {
        border-radius: 0px !important;
      }

      .el-form-item__content .ico {
        position: absolute;
        top: 0px;
        left: 0px;
        z-index: 999;
        width: 40px;
        height: 39px;
        text-align: center;
        border-right: 1px solid #ccc;
      }

      .ewmbox {
        width: 100%;
        height: 240px;
        margin-top: -25px;

        .ewm {
          width: 140px;
          height: 140px;
          margin: 20px auto;

          p {
            font-size: 12px;
            padding-left: 40px;
            margin: 0;
          }
        }

        .ewmicon {
          width: 140px;
          margin: 20px auto 0;

          .iconfont {
            float: left;
          }

          p {
            font-size: 12px;
            padding-left: 40px;
            margin: 0;
          }
        }
      }
    }
  }
</style>