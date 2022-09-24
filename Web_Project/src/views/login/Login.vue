<template>
    <div class = "Login">
        <!-- 背景 -->
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <!-- 内容 -->
        <div class="content">
            <table border="1" style="wight:100%" id="table">
                <!-- 标题 -->
                <tr>
                    <td colspan="3">
                        <h1 v-if=msg>
                            {{msg}}
                        </h1>
                        <h1 v-else>
                            请先登录以查看数据
                        </h1>
                    </td>
                </tr>
                <!-- 用户名 -->
                <tr>
                    <td align="center" style = "width: 20%;">
                        <img src="../../assets/login/login/userName.png" width="40%">
                    </td>
                    <td colspan="2" align="left">
                        <input v-model="userName" placeholder="请输入用户名" id="input" /> (11)
                    </td>
                </tr>
                <!-- 密码 -->
                <tr>
                    <td align="center" style = "width: 20%;">
                        <img src="../../assets/login/login/password.png" width="40%">
                    </td>
                    <td colspan="2" align="left">
                        <input v-model="password" type="password" placeholder="请输入密码" id="input"/>(22)
                    </td>
                </tr>
                <!-- 按键 -->
                <tr>
                    <td colspan="3" align="center" id="submit"> 
                        <button @click="login()" id="button">登录</button>
                    </td>
                </tr>
            </table>
            
        </div>
    </div>
    

</template>

<script>
    import axios from "axios"
    // axios.defaults.withCredentials = true
    import Cookies from 'js-cookie'

    export default{
        name: "login",
        data() {
            return {
                // 背景
                imgSrcBK:require('../../assets/login/login/background.jpeg'),

                // 内容
                msg: "",
                userName: "LocalHost",
                password: "127.0.0.1"
            }
        },
        methods: {
            login(){
                 if ( !this.userName || !this.password) {
                    this.msg = "请输入用户名或密码"
                } 
                let that = this

                axios.post(
                    "http://127.0.0.1:1432/api/v1/login?" + 
                    "UserName=" + this.userName + "&" + 
                    "Password=" + this.password
                ).then(function(res){
                    if (res.data.Code == 0 ){
                        console.log(res)
                        that.msg = res.data.Data
                        // let session = document.cookie.match(
                        //     "(^|[^;]+)\\s*" + "camp-session" + "\\s*=\\s*([^;]+)"
                        //  );
                        console.log(Cookies.get())
                        // that.$router.push({
                        //     path: `/data/salesVolume`
                        // })
                    } else {
                        that.msg = "账号或密码错误"
                    }
                }).catch(function(err){
                    console.log(err)
                })
            }
        },
        mounted() {
            this.msg = this.$route.query.msg
        }
    }

</script>

<style>

    .background {
        width:100%;
        height:100%; 
        top:0%;
        left: 0%;
        z-index:-1;
        position: absolute;
    }

    /* 内容设置 */
    .content {
        opacity: 0.7; /* 透明度 */
        width: 20%;
        z-index:1;
        position: absolute;
        top: 20%;
        right: 5%;
    }

    #input{
        width: 60%;
    }

    #table{
        color : silver
    }

    #button{
        background-color:darkturquoise;
        color: aliceblue;
    }
</style>
