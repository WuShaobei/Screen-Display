<template>
    <div>
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="login">
        
            <div class="content">
            
                <div class="leftView">
                    <div v-if="todo==='tips'" class="leftValue">
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
                    <!-- TODO 注册 -->
          
                </div>

                <div class="rightView">
                    <div class="rightValue">
                        <div class="box1">
                            <div class="text">
                                {{
                                    msg ? msg : "登录页"
                                }}
                            </div>
                        </div>
                        <div class="box2">
                            <div class="leftBox"><div class="text">用户名:</div></div>
                            <div class="rightBox"><div class="text">
                                <input v-model="Username" placeholder="请输入用户名" id="input" /> 
                            </div></div>
                        </div>
                        <div class="box3">
                            <div class="leftBox"><div class="text">密码</div></div>
                            <div class="rightBox"><div class="text">
                                <input v-model="Password" type="password" placeholder="请输入密码" id="input"/>
                            </div></div>
                        </div>
                        <div class="box4">
                            <div class="leftBox"><div class="text"></div></div>
                            <div class="rightBox"><div class="text">
                                <input type="radio" v-model="setCookie" value="true" id="true"/> 30天免登录
                            </div></div>
                        </div>
                        <div class="box5">

                            <div class="leftBox"><div class="text">
                                <!-- TODO 注册 -->
                                <button v-if="todo==='tips'"  @click="toClose()" id="button">关闭提示</button>
                                <button v-else @click="toTips()" id="button">查看提示</button>
                            </div></div>
                            
                            <div class="rightBox"><div class="text">
                                <button @click="login()" id="button">登录</button>
                            </div></div>
                       
                        </div>

                    </div>
                </div>
                    
            </div>


            <div class="indexes">
                <div class="text">
                    <h1 id="h1">
                        中<br />
                        国<br />
                        餐<br />
                        饮<br />
                        市<br />
                        场
                    </h1>
                </div>
            </div>
        
        </div>
    </div>
    

</template>

<script>
    
    import axios from "axios"

    export default{
        name: "login",

        data(){
            return{
                // 背景
                imgSrcBK:require('../../assets/login/login/background.jpeg'),
                msg:"",
                todo:"tips",
                setCookie: "",
                SessionKey: "",
                Username: "",
                Password: "",
                RealName:"",
                Identity:"",
            }
        },
        methods:{
            toClose(){
                this.todo = "register"
            },
            toTips(){
                this.todo = "tips"
            },login(){
                 if ( !this.userName || !this.password) {
                    this.msg = "请输入用户名或密码"
                } 
                let that = this

                axios.post(
                    "http://127.0.0.1:1432/api/v1/login?",{                        
                        Username : this.Username, 
                        Password : this.Password,
                    }
                ).then(function(res){
                    if (res.data.Code == 0 ){
                        console.log(res)
                        if (that.setCookie == 'true'){
                            that.$cookies.set("camp-session", res.data.Session)                  
                        }
                        that.Username = res.data.Data.Username
                        that.RealName = res.data.Data.RealName
                        that.Identity = res.data.Data.Identity
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
                    path: `/data/data`,
                    query: {
                        Username: this.Username,
                        RealName: this.RealName,
                        Identity: this.Identity
                    },
                })
            }
        },
        mounted() {
            let that = this
            let SessionKey = this.$cookies.get("camp-session")
            console.log(SessionKey)
            axios.post(
                "http://127.0.0.1:1432/api/v1/login?",{                        
                    SessionKey: SessionKey
                }
            ).then(function(res){
                if (res.data.Code == 0 ){
                    that.Username = res.data.Data.Username
                    that.RealName = res.data.Data.RealName
                    that.Identity = res.data.Data.Identity
                    that.jumpTo()
                } else {
                    that.msg = ""
                }
            }).catch(function(err){
                console.log(err)
            })
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
            width:80%;
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
                    position: absolute;
                }
        .indexes{
            height:100%;
            width:20%;
            display: flex;
            position: absolute;
            left: 80%;
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
    background-color: aqua;
    text-align: center;
}
.box2{
    height: 20%;
    width: 100%;
    top:20%;
    position: absolute;
    display: flex;
    background-color: aqua;
}
.box3{
    height: 20%;
    width: 100%;
    top:40%;
    position: absolute;
    display: flex;
    background-color: aqua;
}
.box4{
    height: 20%;
    width: 100%;
    top:60%;
    position: absolute;
    display: flex;
    background-color: aqua;
}
.box5{
    height: 20%;
    width: 100%;
    top:80%;
    position: absolute;
    display: flex;
    background-color: aqua;
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
    height: 100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
    background-color: aquamarine;
}

 #input{
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