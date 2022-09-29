<!--
 * @Author: WuShaobei
 * @Date: 2022-09-28 09:21:46
 * @LastEditTime: 2022-09-29 14:50:23
 * @FilePath: /Web_Project/src/views/login/Login.vue
 * @Description: 登录页
-->

<template>
    <div>
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="login">
        
            <div class="content">
            
                <div class="leftView">
                    <div v-if="todo==='login'" class="leftValue">
                        <div class="box1">
                            <div class="tableBox1"><div class="text">Id</div></div>
                            <div class="tableBox2"><div class="text">用户名</div></div>
                            <div class="tableBox3"><div class="text">密码</div></div>
                            <div class="tableBox4"><div class="text">类型</div></div>
                        </div>
                        <div class="box2">
                            <div class="tableBox1"><div class="text">1</div></div>
                            <div class="tableBox2"><div class="text">LocalHost</div></div>
                            <div class="tableBox3"><div class="text">127.0.0.1</div></div>
                            <div class="tableBox4"><div class="text">管理员</div></div>
                        </div>
                        <div class="box3">
                            <div class="tableBox1"><div class="text">2</div></div>
                            <div class="tableBox2"><div class="text">Investor</div></div>
                            <div class="tableBox3"><div class="text">Investor</div></div>
                            <div class="tableBox4"><div class="text">投资人</div></div>
                        </div>
                        <div class="box4">
                            <div class="tableBox1"><div class="text">3</div></div>
                            <div class="tableBox2"><div class="text">Practitioner</div></div>
                            <div class="tableBox3"><div class="text">Practitioner</div></div>
                            <div class="tableBox4"><div class="text">从业者</div></div>
                        </div>
                        <div class="box5">
                            <div class="tableBox1"><div class="text">4</div></div>
                            <div class="tableBox2"><div class="text">Tourist</div></div>
                            <div class="tableBox3"><div class="text">Tourist</div></div>
                            <div class="tableBox4"><div class="text">游客</div></div>
                        </div>
                    </div>  
                </div>

                <div class="rightView">
                    <div v-if="todo==='register'" class="rightValue">
                        <div class="box1">
                            <div class="text">
                                {{
                                    "注册页"
                                }}
                            </div>
                        </div>
                        <div class="box2">
                            <div class="leftBox"><div class="text">用户名:</div></div>
                            <div class="rightBox"><div class="text">
                                <input v-model="Username" placeholder="请输入用户名" id="inputUsername" /> 
                            </div></div>
                        </div>
                        <div class="box3">
                            <div class="leftBox"><div class="text">密码</div></div>
                            <div class="rightBox"><div class="text">
                                <input v-model="Password" type="text" placeholder="请输入密码" id="inputPassword"/>
                            </div></div>
                        </div>
                        <div class="box4">
                            <div class="leftBox"><div class="text">真实姓名</div></div>
                            <div class="rightBox"><div class="text">
                                <input v-model="RealName" type="text" placeholder="请输入真实姓名" id="inputRealName"/>
                            </div></div>
                        </div>
                        <div class="box5">
                            <div class="leftBox"><div class="text">
                                <button @click="toLogin()" id="button">登录</button>
                            </div></div>
                            <div class="rightBox"><div class="text">
                                <button @click="register()" id="button">注册</button>
                            </div></div>
                        
                        </div>
                    </div>
                    <div v-else class="rightValue">
                        <div class="box1">
                            <div class="text">
                                {{
                                    "登录页"
                                }}
                            </div>
                        </div>
                        <div class="box2">
                            <div class="leftBox"><div class="text">用户名:</div></div>
                            <div class="rightBox"><div class="text">
                                <input v-model="Username" placeholder="请输入用户名" id="inputUsername" /> 
                            </div></div>
                        </div>
                        <div class="box3">
                            <div class="leftBox"><div class="text">密码</div></div>
                            <div class="rightBox"><div class="text">
                                <input v-model="Password" type="password" placeholder="请输入密码" id="inputPassword"/>
                            </div></div>
                        </div>
                        <div class="box4">
                            <div class="leftBox"><div class="text"></div></div>
                            <div class="rightBox"><div class="text">
                                <input type="radio" v-model="setCookie" value="true" id="true"/> 3天免登录
                            </div></div>
                        </div>
                        <div class="box5">

                            <div class="leftBox"><div class="text">
                                <button  @click="toRegister()" id="button">注册账号</button>
                            </div></div>

                            <div class="rightBox"><div class="text">
                                <button @click="login()" id="button">登录</button>
                            </div></div>
                    
                        </div>

                    </div>
                </div>
                    
            </div>


            <div class="Title">
                <div class="text">
                    <h3 id="h1">
                        中<br />
                        国<br />
                        餐<br />
                        饮<br />
                        市<br />
                        场
                    </h3>
                </div>
            </div>
        
        </div>
    </div>
    

</template>

<script>
    
    import * as PostApiUsers from "../../Web/PostApiUsers"
    export default{
        name: "login",

        data(){
            return{
                // 背景
                imgSrcBK:require('../../assets/background.jpeg'),
                todo:"login",
                setCookie: "",
                SessionKey: "",
                Username: "",
                Password: "",
                RealName:"",
                Identity:"3",
                Id:"",
            }
        },
        methods:{
            clearData(){
                this.Username = ""
                this.Password = ""
                this.RealName = ""
            },
            toRegister(){
                this.todo = "register"
                this.clearData()
            },
            toLogin(){
                this.todo = "login"
                this.clearData()
            },
            login(){
                if ( this.Username=="" || this.Password=="") {
                    alert("请输入用户名或密码")
                    return 
                } 
                PostApiUsers.postApiLoginByPassword(this.Username, this.Password, 
                    (res) => {
                        if ( ! res ) {
                            alert("账号或密码错误")
                        } else {
                            if (this.setCookie == 'true') {
                                this.$cookies.set("camp-session", res.SessionKey)    
                            }
                            this.jumpTo( res.Id )
                        }
                    }
                )
            },
            register(){
                if ( this.Username=="" || this.Password=="" || this.RealName=="") {
                    alert("请输入完整信息")
                    return
                }
                PostApiUsers.postApiRegister(this.Username, this.Password, this.RealName, this.Identity,
                    (res) => {
                        alert(res)
                        if ( res == "注册成功" ) {
                            this.todo = "login"
                        }
                    }
                )
                
            },
            jumpTo( Id ){
                this.$router.push({
                    path: `/data/Screen`,
                    query: {
                        Id: Id
                    },
                })
            }
        },
        mounted() {
            // 免登录
            let SessionKey = this.$cookies.get("camp-session")
            if (SessionKey) {
               PostApiUsers.postApiLoginBySession(SessionKey, (res) => {
                    if (res) {
                        this.jumpTo(res.Id)
                    }
               })
            }
        }
        
    }

</script>


<style>
    .background{
        height:100%; 
        width:100%;
        top: 0%;
        left: 0%;
        z-index: -1;
        position: absolute;
    }

    .login{
        height: 100%;
        width: 100%;
        top: 0%;
        left: 0%;
        z-index: 1;
        position: absolute;
        display: flex;
        opacity: 0.7; 
    }    
        


        .content{
            height:100%; 
            width:95%;
            display: flex;
            position: absolute;     
        }

            
            .leftView{
                height:100%; 
                width:50%;
                left: 0%;
                display: flex;
                position: absolute;
            }

                .leftValue{
                    height:60%; 
                    width:80%;
                    left: 20%;
                    top:20%;
                    display: flex;    
                    flex-direction: column;
                    background-color: aliceblue;
                    position: absolute
                }

        
            .rightView{
                height:100%; 
                width:50%;
                left: 50%;
                display: flex;    
                position: absolute;     
            }

                .rightValue{
                    height:60%; 
                    width:80%;
                    left: 0%;
                    top:20%;
                    display: flex;    
                    background-color: aliceblue;
                    position: absolute;
                }
        .Title{
            height:100%;
            width:5%;
            display: flex;
            position: absolute;
            left: 95%;
            background-color: rgb(124, 94, 16);
            text-align: center;
        }

/* box */
.box1{
    height: 20%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    text-align: center;
}
.box2{
    height: 20%;
    width: 100%;
    top:20%;
    position: absolute;
    display: flex;
}
.box3{
    height: 20%;
    width: 100%;
    top:40%;
    position: absolute;
    display: flex;
}
.box4{
    height: 20%;
    width: 100%;
    top:60%;
    position: absolute;
    display: flex;
}
.box5{
    height: 20%;
    width: 100%;
    top:80%;
    position: absolute;
    display: flex;
}
.tableBox1{
    height: 100%;
    width: 25%;
    left: 0%;
    position: absolute;
}
.tableBox2{
    height: 100%;
    width: 25%;
    left: 25%;
    position: absolute;
}
.tableBox3{
    height: 100%;
    width: 25%;
    left: 50%;
    position: absolute;
}
.tableBox4{
    height: 100%;
    width: 25%;
    left: 75%;
    position: absolute;
}
.leftBox{
    height: 100%;
    width: 30%;
    position: absolute;

}
.rightBox{
    height: 100%;
    width: 70%;
    left: 30%;
    position: absolute;
}

/* text */
.text{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
}

 #inputUsername{
    text-align: left;
    height: 80%;
    width: 80%;
    top: 10%; 
    left: 0%;
    position: absolute;
 }
 #inputPassword{
    text-align: left;
    height: 80%;
    width: 80%;
    top: 10%; 
    left: 0%;
    position: absolute;
 }
 #inputRealName{
    text-align: left;
    height: 80%;
    width: 80%;
    top: 10%; 
    left: 0%;
    position: absolute;
 }

 #h1{
    text-align: center;
    height: 20%;
    width: 100%;
    /* background-color: aqua; */  
 }
</style>