<!--
 * @Date: 2022-10-09 16:49:59
 * @LastEditTime: 2022-10-14 17:58:37
 * @FilePath: /frontEnd/src/views/login.vue
 * @Description: 登录页
-->
<template>
    <div class = "viewLogin">
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="content">
<!------------------------------------------------------------------------------>
<!------------------------------------ 登录页 ----------------------------------->
<!------------------------------------------------------------------------------>
            <div v-if="nowIs==='login'" class="dataTable" style="display:flex;flex-direction: column;">
                <div class="tableTitle"><h1>登录页</h1></div>
                <div class="tableTr">
                    <div  class="tableTdLeft">用户名</div>
                    <div  class="tableTdRight"><input v-model="username" placeholder="请输入用户名"/></div>
                </div>
                <div class="tableTr"><form>
                    <div  class="tableTdLeft">密码</div>
                    <div  class="tableTdRight"><input v-model="password" type="password" placeholder="请输入密码" autocomplete/></div>
                </form></div>
                <div class="tableTr">
                    <div style="left: 50%; position: absolute;"><input type="checkbox" v-model="setCookie" value="true" id="true"/> 3天免登录</div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft"><button  @click="changeState('register')" id="button">注册</button></div>
                    <div  class="tableTdRight"><button @click="login()" id="button">登录</button></div>
                    <br />
                </div>
            </div>
<!------------------------------------------------------------------------------>
<!------------------------------------ 注册页 ----------------------------------->
<!------------------------------------------------------------------------------>
            <div v-if="nowIs==='register'" class="dataTable" style="display:flex;flex-direction: column;">
                <div class="tableTitle"><h1>注册页</h1></div>
                <div class="tableTr">
                    <div  class="tableTdLeft">用户名</div>
                    <div  class="tableTdRight"><input v-model="username" placeholder="请输入用户名" /></div>
                </div>
                <div class="tableTr"><form>
                    <div  class="tableTdLeft">密码</div>
                    <div  class="tableTdRight"><input v-model="password" type="password" placeholder="请输入密码" autocomplete/></div>
                </form></div>
                <div class="tableTr">
                    <div  class="tableTdLeft">真实姓名</div>
                    <div  class="tableTdRight"><input v-model="realName" placeholder="请输入真实姓名" /></div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft">用户类型</div>
                    <div  class="tableTdRight">
                        <select v-model="identity" value="3" >
                            <option value="1">投资人</option>
                            <option value="2">从业者</option>
                            <option value="3" selected>游客</option>
                        </select>
                    </div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft"><button  @click="changeState('login')" id="button">登录</button></div>
                    <div  class="tableTdRight"><button @click="register()" id="button">注册</button></div>
                    <br/>
                </div>
            </div>
        </div>

        <div class="title" style="display: flex;flex-direction: column;">
            <div v-for="c in '中国餐饮大数据平台'" style="width:100%;">
                <h2>{{c}}</h2>
            </div>
        </div>
    </div>
</template>

<script>

    import * as userApi from '../api/user'


    export default{
        data () {
            return {
                // 背景
                imgSrcBK:require('../assets/background.jpeg'),
                nowIs: "login",
                username : "",
                password : "",
                setCookie : "",
                realName : "",
                identity : "",
            }
        },
        methods : {
            /**
             * @description: 切换显示的界面  
             * @param {*} todo : 将显示的界面  login / register
             * @return {*}
             */            
            changeState( todo ) {
                this.nowIs = todo
                this.username = ""
                this.password = ""
                this.realName = ""
                this.identity = ""
            },

            /**
             * @description: 基础参数校验 登录
             * @return {*}
             */            
            login(){
                if ( this.username=="" ) { alert("请输入用户名"); return }
                if ( this.password=="" ) { alert("请输入密码"); return }
                if ( this.username.length <  8 ||  this.username.length > 15 ) {alert("用户名或密码错误"); return }
                if ( this.password.length <  6 ) {alert("用户名或密码错误"); return }
                userApi.postLoginByPasswordApi(this.username, this.password, this.setCookie, 
                    (res) => {
                        if ( res.code != 0 ) {
                            alert("用户名或密码错误")
                        } else {
                            if (this.setCookie != "") {
                                this.$cookies.set("camp-session", res.data.sessionKey)    
                                console.log(this.$cookies.get("camp-session"))
                            }
                            this.jumpTo( res.data.id )
                        }
                    }
                )
            },
            
            /**
             * @description: 携带用户 id 跳转到数据页
             * @param {*} id 用户 id
             * @return {*}
             */            
            jumpTo( id ) {
                this.$router.push({
                    path: `/display`,
                    query: {
                        id: id
                    },
                })
            },

            /**
             * @description: 基础参数校验 登录 
             * @return {*}
             */            
            register(){
                if ( this.username=="" ) { alert("请输入用户名"); return }
                if ( this.username.length <  8 ||  this.username.length > 15 ){alert("请输入长度为 8 到 15 的用户名"); return }
                if ( this.password=="" ) { alert("请输入密码"); return }
                if ( this.password.length <  6 ) {alert("请输入长度为 6 位及以上的密码"); return }
                if ( this.realName=="" ) { alert("请输入真实姓名"); return }
                if ( this.identity=="" ) { alert("请选择用户类型"); return }
                
                
                userApi.postRegisterApi(this.username, this.password, this.realName,this.identity,
                    (res) => {
                        if ( res.code == 0 ) {
                            alert("注册成功 id : " + res.data)
                            this.nowIs = "login"
                        } else {
                            alert("用户名已存在")
                        }
                    }
                )
                
            },
        },
        mounted() {
            // 尝试通过 session 值登录
            let sessionKey = this.$cookies.get("camp-session")
            if (sessionKey) {
                userApi.postLoginBySessionKeyApi(sessionKey, (res) => {
                    if (res.code == 0) {
                        this.jumpTo(res.data.id)
                    }
               })
            }
        }
    }
</script>

<style>
/* ========================================================= */
/* =======================底层界面=========================== */
/* ========================================================= */
.viewLogin{

    height: 100%; top: 0%;
    width: 100%; left: 0%;
    position: absolute;

    /* 缩放正常显示 */
    min-width: 1400px;
    min-height: 800px;
    overflow: hidden;
}

/* 背景设置 */
.background{
    height: 100%; top: 0%;
    width: 100%; left: 0%;
    z-index: -1;
    position: absolute;
}

.content{
    height: 100%; top: 0%;
    width: 95%; left: 0%; 
    position: absolute;
    opacity: 0.7;
    z-index: 1;
}

.title{
    height: 100%; top: 0%;
    width: 5%; left: 95%;
    position: absolute;
    opacity: 0.7;
    z-index: 1;
    background-color: rgb(124, 94, 16);
}

/* ========================================================= */
/* ========================数据页============================ */
/* ========================================================= */

.dataTable{
    height: auto; top: 20%;
    width: auto; left: 60%;
    min-width: 25%;
    position: absolute;
    background-color: aliceblue;
}

.tableTitle{
    width: 100%;
}

.tableTr{
    width: 100%;
    margin-top: 5%;
    margin-bottom: 5%;
    display: flex;
    flex-direction: row;
}
.tableTdLeft{
    width: 40%; left: 0;
    position: absolute;
}
.tableTdRight{
    width: 60%; left: 40%;
    position: absolute;
}

/* ========================================================= */
/* ==========================其他============================ */
/* ========================================================= */

#button{
    width: 40%;
    background-color: bisque;
}

</style>