<!--
 * @Author: WuShaobei
 * @Date: 2022-10-09 16:49:59
 * @LastEditTime: 2022-10-09 21:03:44
 * @FilePath: /Screen-Display/v2/frontEnd/src/views/login.vue
 * @Description: 登录页
-->
<template>
    <div class = "viewLogin">
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="content">
            <div v-if="nowIs==='login'" class="dataTable" style="display:flex;flex-direction: column;">
                <div class="tableTitle"><h1>登录页</h1></div>
                <div class="tableTr">
                    <div  class="tableTdLeft">用户名</div>
                    <div  class="tableTdRight"><input v-model="username" placeholder="请输入用户名" /></div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft">密码</div>
                    <div  class="tableTdRight"><input v-model="password" type="password" placeholder="请输入密码" /></div>
                </div>
                <div class="tableTr">
                    <div style="left: 50%; position: absolute;"><input type="radio" v-model="setCookie" value="true" id="true"/> 3天免登录</div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft"><button  @click="changeState('register')" id="button">注册</button></div>
                    <div  class="tableTdRight"><button @click="login()" id="button">登录</button></div>
                    <br />
                </div>
                <!-- TODO 删除 -->
                <div class="tableTr" style="text-align: left;">
                    Tip: <br />默认用户名 : admin<br /> 密码 admin
                </div>
            </div>

            <div v-if="nowIs==='register'" class="dataTable" style="display:flex;flex-direction: column;">
                <div class="tableTitle"><h1>注册页</h1></div>
                <div class="tableTr">
                    <div  class="tableTdLeft">用户名</div>
                    <div  class="tableTdRight"><input v-model="username" placeholder="请输入用户名" /></div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft">密码</div>
                    <div  class="tableTdRight"><input v-model="password" type="password" placeholder="请输入密码" /></div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft">真实姓名</div>
                    <div  class="tableTdRight"><input v-model="realName" type="password" placeholder="请输入真实姓名" /></div>
                </div>
                <div class="tableTr">
                    <div  class="tableTdLeft"><button  @click="changeState('login')" id="button">登录</button></div>
                    <div  class="tableTdRight"><button @click="login()" id="button">注册</button></div>
                    <br/>
                </div>
            </div>
        </div>

        <div class="title" style="display: flex;flex-direction: column;">
            <div v-for="c in '中国餐饮市场'" style="width: ;100%">
                <h1>{{c}}</h1>
            </div>
        </div>
    </div>
</template>

<script>
    export default{
        data () {
            return {
                // 背景
                imgSrcBK:require('../assets/background.jpeg'),
                nowIs: "login",
                username : "",
                password : "",
                realName : "",
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
            }
        }
    }
</script>

<style>
/* ========================================================= */
/* =======================底层界面=========================== */
/* ========================================================= */
.viewLogin{
    height: 100%;
    width: 100%;
    top: 0%;
    left: 0%;

    /* 缩放正常显示 */
    min-width: 1000px;
    min-height: 600px;
    overflow: hidden;

    position: absolute;
}

    /* 背景设置 */
    .background{
        height: 100%;
        width: 100%;
        left: 0%;
        right: 0%;
        z-index: -1;
        position: absolute;
    }

    .content{
        opacity: 0.7;
        height: 100%;
        width: 95%;
        left: 0%;
        right: 0%;
        z-index: 1;
        position: absolute;
    }
        
    .title{
        height: 100%;
        width: 5%;
        left: 95%;
        right: 0%;
        z-index: 1;
        position: absolute;
        opacity: 0.7;
        background-color: rgb(124, 94, 16);
    }

/* ========================================================= */
/* ========================数据页============================ */
/* ========================================================= */

.dataTable{
    height: auto;
    width: auto;
    min-width: 25%;
    top: 20%;
    left: 60%;
    position: absolute;
    background-color: aliceblue;
}

.tableTitle{
    width: 100%;
}
.tableTr{
    margin-top: 5%;
    margin-bottom: 5%;
    width: 100%;
    display: flex;
    flex-direction: row;
}
.tableTdLeft{
    left: 0;
    width: 40%;
    position: absolute;
}
.tableTdRight{
    width: 60%;
    left: 40%;
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