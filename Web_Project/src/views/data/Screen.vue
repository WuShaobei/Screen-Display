<template>
    <div class="AllDataView">
        
        <!-- BackGround -->
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <!-- APi Title -->
        <div class="dataView">
        
            <!-- Api -->
            <div class="content">
                
                <!-- Api4 -->
                <div class="leftView">
                    <e-charts :option="Api4Op"/>
                </div>
                
                <!-- Api1 & Api5 -->
                <div class="centerView">
                    
                    <!-- Api1 -->
                    <div class="top1">
                        <e-charts :option="Api1Op"/>
                    </div>
                    
                    <!-- Api5 -->
                    <div class="down1">
                        <div id="scroll_in2" class="htp-list_scroll_in" style="height: 100%;overflow-y:hidden;font-size:20px;">
                            
                            <table>
                                <tr>
                                    <td colspan="3">
                                        <h1>
                                            2021年-2022年预制菜行业融资情况
                                        </h1>
                                    </td>
                                </tr>
                                <tr>
                                    <td style="height: 20%;"> <h3>品牌</h3> </td>
                                    <td> <h3>时间</h3> </td>
                                    <td> <h3> 轮次</h3> </td>
                                    <td> <h3>投资商</h3> </td>
                                </tr>
                                <tr v-for="data in Api5Value">
                                    <td style="height: 20%;">
                                        {{data.Brand}}
                                    </td>
                                    <td>
                                        {{data.Time}}
                                    </td>
                                    <td>
                                        {{data.FundingRound}}
                                    </td>
                                    <td>
                                        {{data.Investor}}
                                    </td>
                                </tr>
                            </table>

                        </div>
                    </div>
                
                </div>

                <!-- ApiUser & Api3 & Api2 -->
                <div class="rightView">

                    <!-- ApiUser -->
                    <div class="top0">    
                        <div class="table0Box1">
                            <div class="text0">
                                <h3>
                                    欢迎{{getIdentity()}} {{whoAmI.RealName}}
                                </h3>
                            </div>
                        </div>
                        <div class="table0Box2">
                            <div class="text0">
                                Id : {{whoAmI.Id}}
                                </div>
                        </div>
                        <div class="table0Box3">
                            <div class="text0">
                                用户名 : {{whoAmI.Username}}
                            </div>
                        </div>
                    </div>

                    <!-- Api3 -->
                    <div class="top2">
                        <div class="table1" style="font-size:20px;">
                            <div class="table1Box" >
                                <div class="text" style="text-align: left;">2018年-2021年餐饮就业人数行业平均报酬</div>
                            </div>
                            <div class="table1Box">
                                <div class="table1BoxData1"><div class="text">时间</div></div>
                                <div class="table1BoxData2"><div class="text">平均报酬</div></div>
                                <div class="table1BoxData3"><div class="text">就业人数</div></div>
                            </div>
                            <div v-for="data in Api3Value" class="table1Box">
                                <div class="table1BoxData1"><div class="text">{{data.Year}}</div></div>
                                <div class="table1BoxData2"><div class="text">{{data.AvgSalary}}</div></div>
                                <div class="table1BoxData3"><div class="text">{{data.CountSalary}}</div></div>
                            </div>
                        </div>
                        
                    </div>

                    <!-- Api2 -->
                    <div class="down2">    
                        <e-charts :option="Api2Op"/>
                    </div>

                </div>
                    
            </div>

            <!-- Title -->
            <div class="Title">

                <div class="text">

                    <h3>
                        中<br />
                        国<br />
                        餐<br />
                        饮<br />
                        数<br />
                        据<br />
                        
                        
                        <button @click="logout">注销</button>
                        
                    </h3>
                </div>
                
            </div>
        
        </div>
    </div>
    

</template>

<script>
    
    import axios from 'axios'
    import * as ApiUser from '../../Web/ApiUser'
    import {getApi1Data} from '../../Web/Api1'
    import {getApi2Data} from '../../Web/Api2'
    import {getApi3Data} from '../../Web/Api3'
    import {getApi4Data} from '../../Web/Api4'
    import {getApi5Data} from '../../Web/Api5'
    export default{
        name: "login",

        data(){
            return{
                // 背景
                imgSrcBK:require('../../assets/background.jpeg'),
                Api1Op : {},
                Api2Op : {},
                Api3Value:{},
                Api4Op : {},
                Api5Value:{},
                whoAmI:{},
            }
        },
        methods:{
            getIdentity(){
                if (this.whoAmI.Identity == '0') { return "管理员"}
                if (this.whoAmI.Identity == '1') { return "投资方"}
                if (this.whoAmI.Identity == '2') { return "从业者"}
                if (this.whoAmI.Identity == '3') { return "游客"}
            },
            logout(){
                this.$cookies.set("camp-session", "", -1)
                
                this.$router.push({
                    path: `/`
                })
                
            },
            Api1(){
                getApi1Data((res) =>{
                    this.Api1Op = res.Data
                })
            },
            Api2(){
                getApi2Data((res) =>{
                    this.Api2Op = res.Data
                })
            },
            Api3(){
                getApi3Data((res) =>{
                    this.Api3Value = res.Data
                })
            },
            Api4(){
                getApi4Data((res) =>{
                    this.Api4Op = res.Data
                })
                
            },
            Api5(){
                getApi5Data((res) =>{
                    this.Api5Value = res.Data
                })
            },    // 添加自动滚动
            /* 
            Id   需要滚动的区域 id名称
            */
            autoSroll(Id) {
            // flag 为true时停止滚动
                var flag = false;
                // 定时器
                var timer;
                function roll() {
                    var h = -1;
                    timer = setInterval(function() {
                        flag = true;
                        if (flag) {
                            var current = document.getElementById(Id).scrollTop;
                            if (current == h) {
                                //滚动到底端,返回顶端
                                h = 0;
                                document.getElementById(Id).scrollTop = h + 1;
                            } else {
                                //以25ms/3.5px的速度滚动
                                h = current;
                                document.getElementById(Id).scrollTop = h + 1;
                            }
                            
                        }
                        //获取当前滚动条高度
                        }, 50);
                    }
                // setTimeout(function() {
                    //滚动区域内单击鼠标 滚动暂停 再次点击鼠标 继续滚动
                    document.getElementById(Id).onclick = () => {
                    
                    if (flag) {
                        flag = false;
                        clearInterval(timer);
                    } else {
                        roll();
                    }
                    };
                    roll();
                // }, 2000);
                },    
        },
        mounted(){
            ApiUser.getWhoAmIData(this.$route.query.Id , (res)=>{
                this.whoAmI = res.Data
            })
            this.Api1()
            this.Api2()
            this.Api3()
            this.Api4()
            this.Api5()
            this.autoSroll("scroll_in2")
            let that = this
            axios.post(
                "http://127.0.0.1:1432/api/v1/whoAmI?",{                        
                    Id : that.Id
                }
            ).then(function(res){
                if (res.data.Code == 0 ){
                    that.Username = res.data.Data.Username
                    that.RealName = res.data.Data.RealName
                    that.Identity = res.data.Data.Identity
                } else {
                    
                }
            }).catch(function(err){
                console.log(err)
            })

        },
    }

</script>


<style>

 .AllDataView{
    height:100%; 
    width:100%;
    top: 0%;
    left: 0%;
    position: absolute;
 }   
    .background{
        height:100%; 
        width:100%;
        top: 0%;
        left: 0%;
        z-index: -1;
        position: absolute;
    }

    .dataView{
        height: 100%;
        width: 100%;
        top: 0%;
        left: 0%;
        z-index: 1;
        position: absolute;
        display: flex;
        
        opacity: 0.750; 
        /* background-color: antiquewhite; */
    }    

        .content{
            height:100%; 
            width:95%;
            display: flex;
            position: absolute;     
            /* background-color: black;    */
        }

            .leftView{
                height:100%; 
                width:25%;
                left: 0%;
                display: flex;    
                position: absolute;
                background-color: rgb(255, 255, 255);
            }
            .centerView{
                height:100%; 
                width:45%;
                left: 25%;
                display: flex;    
                position: absolute;
                background-color: aliceblue;
            }      
            .rightView{
                height:100%; 
                width:30%;
                left: 70%;
                display: flex;    
                position: absolute;        
                background-color: azure;
            }

        .Title{
            height:100%;
            width:5%;
            top: 0%;
            display: flex;
            position: absolute;
            left: 95%;
            background-color: rgb(124, 94, 16);
        }


/* text */
.text{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
    overflow-y:hidden
}
.text0{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 25px; 
    text-align: center;  
    overflow-y:hidden
}
.textIndex{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: left;  
}
.top0{
    height: 15%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    flex-direction: row;
    background-color: rgb(240,240,240)
}
.top1{
    height: 50%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    background-color: rgb(250,250,250)
}
.top2{
    height: 35%;
    width: 100%;
    top:15%;
    position: absolute;
    display: flex;
    background-color:rgb(245, 245, 245);
}
.down1{
    height: 50%;
    width: 100%;
    top:50%;
    position: absolute;
    display: flex;
    background-color:rgb(245, 245, 245);
}
.down2{
    height: 50%;
    width: 100%;
    top:50%;
    position: absolute;
    display: flex;
    background-color: rgb(250,250,250);
}

.htp-list_scroll_in{
    height: 100%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
}

.table1{
    height: 100%;
    width: 100%;
    position: absolute;
    top: 0%;
    left: 0%;
    display: flex;
    flex-direction:column;
}
.table1Box{
    width: 100%;
    height: 15%;
    top: auto;
    display: flex;
    flex-direction: column;
}
.table0Box1{
    width: 100%;
    height: 50%;
    top:0%;
    display: flex;
    position: absolute;
}
.table0Box2{
    width: 100%;
    height: 25%;
    top: 50%;
    position: absolute;
    display: flex;
}
.table0Box3{
    height: 25%;
    width: 100%;
    top: 75%;
    position: absolute;
    display: flex;
}
.table1BoxData1{
    width: 30%;
    height: 100%;
    left: 0%;
    position: absolute;
    display: flex;
}
.table1BoxData2{
    width: 40%;
    height: 100%;
    left: 30%;
    position: absolute;
    display: flex;
}
.table1BoxData3{
    width: 30%;
    height: 100%;
    left: 70%;
    position: absolute;
    display: flex;
}


 #input{
    text-align: left;
    height: 80%;
    width: 80%;
    top: 10%; 
    left: 0%;
    position: absolute;
 }
 
</style>

