<template>
    <div class = "Login">
        <!-- 背景 -->
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>



        <!-- 内容 -->
        <div class="content">
            <div class="dataLine">
                <input v-model="UserName" placeholder="请输入用户名" id="input" /> 
            </div>   
            
            <div class="dataLine">
                <input v-model="Password" type="password" placeholder="请输入密码" id="input"/>
            </div>
            
            
            <div class="dataLine">
                <input type="radio" v-model="setCookie" value="true" id="true"/> 30天免登录
            </div>
            
            
            <div class="dataLine">
                <button @click="login()" id="button">登录</button>
            </div>
            
        </div>


        <div class="contentTip">
            <table border="1" style="wight:100%" id="table">
                <tr style="wight:100%">
                    <td style = "width: 25%;">
                        <b><i>Id</i></b>
                    </td>
                    <td style = "width: 25%;">
                        <b><i>Username</i></b>
                    </td >
                    <td style = "width: 25%;">
                        <b><i>Password</i></b>
                    </td>
                    <td style = "width: 25%;">
                        <b><i>Identity</i></b>
                    </td>
                </tr>
                <tr>
                    <td>
                        1
                    </td>
                    <td>
                        LocalHost
                    </td>
                    <td>
                        127.0.0.1
                    </td>
                    <td>
                        管理员
                    </td>
                </tr>

                <tr>
                    <td>
                        2
                    </td>
                    <td>
                        Investor
                    </td>
                    <td>
                        Investor
                    </td>
                    <td>
                        投资方
                    </td>
                </tr>

                <tr>
                    <td>
                        3
                    </td>
                    <td>
                        Practitioner
                    </td>
                    <td>
                        Practitioner
                    </td>
                    <td>
                        从业者
                    </td>
                </tr>

                <tr>
                    <td>
                        4
                    </td>
                    <td>
                        Tourist
                    </td>
                    <td>
                        Tourist
                    </td>
                    <td>
                        游客
                    </td>
                </tr>
            </table>
        </div>

            
    </div>
    
    

</template>

<script>
    import axios from "axios"
    // axios.defaults.withCredentials = true
    

    export default{
        name: "login",
        data() {
            return {
                // 背景
                imgSrcBK:require('../../assets/login/login/background.jpeg'),

                // 内容
                msg: "",
                UserName: "LocalHost",
                Password: "127.0.0.1",
                setCookie: "",
                SessionKey: ""
                
            }
        },
        methods: {
            login(){
                 if ( !this.userName || !this.password) {
                    this.msg = "请输入用户名或密码"
                } 
                let that = this

                axios.post(
                    "http://127.0.0.1:1432/api/v1/login?",{                        
                        Username : this.userName, 
                        Password : this.password,
                    }
                ).then(function(res){
                    if (res.data.Code == 0 ){
                        console.log(res)
                        that.msg = res.data.Data
                        console.log(that.setCookie)
                        if (that.setCookie == 'true'){
                            that.$cookies.set("camp-session", res.data.Data.Session)                  
                        }
                        that.jumpTo()
                    } else {
                        that.msg = "账号或密码错误"
                    }
                }).catch(function(err){
                    console.log(err)
                })
            },
            jumpTo(){

                console.log("Jump TO")
                this.$router.push({
                    path: `/data/salesVolume`,
                    // query: {
                    //     userName: this.userName,
                    //     password: this.password
                    // },
                })
            }
        },
        mounted() {
            // this.msg = this.$route.query.msg
            let that = this
            let SessionKey = this.$cookies.get("camp-session")
            console.log(SessionKey)
            axios.post(
                "http://127.0.0.1:1432/api/v1/login?",{                        
                    SessionKey: SessionKey
                }
            ).then(function(res){
                console.log(res)
                if (res.data.Code == 0 ){
                    console.log(res)
                    that.msg = res.data.Data
                    console.log(that.setCookie)
                    if (that.setCookie == 'true'){
                        that.$cookies.set("camp-session", res.data.Data.Session)                  
                    }
                    that.jumpTo()
                } else {
                    that.msg = "账号或密码错误"
                }
            }).catch(function(err){
                console.log(err)
            })
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
        display: flex;
        flex-direction: column;
    }

    /* 内容设置 */
    .content {
        opacity: 0.9; /* 透明度 */
        height: 30%;
        width: 30%;
        z-index:1;
        position: absolute;
        display: flex;
        flex-direction: row;
        top: 10%;
        left: 50%;
    }
    .contentTip {
        opacity: 0.9; /* 透明度 */
        height: 30%;
        width: 20%;
        z-index:1;
        position: absolute;
        display: flex;
        top: 10%;
        left: 80%;
    }

    .dataLine{
        opacity: 0.9; /* 透明度 */
        height: 20%;
        width: 100%;
        
        display: flex;
    }
    

    #input{
        width: 60%;
    }

    #table{
        color:rgb(3, 253, 253);
        background-color : antiquewhite;
    }

    #button{
        background-color:darkturquoise;
        color: aliceblue;
    }
</style>
