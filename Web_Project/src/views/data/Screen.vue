<!--
 * @Author: WuShaobei
 * @Date: 2022-09-28 09:21:46
 * @LastEditTime: 2022-09-29 14:56:41
 * @FilePath: /Web_Project/src/views/data/Screen.vue
 * @Description: 大屏显示页
-->

<template>
    <div class="AllDataView">
        
        <!-- BackGround -->
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="dataView">
        

            <div class="content">

                <!-- ChineseCateringBrandStatisticOpData -->
                <div class="leftView" style="overflow-y:hidden;overflow-x:hidden">
                    <e-charts :option="ChineseCateringBrandStatisticOpData"/>
                </div>
                
                <div class="centerView">
                    
                    <!-- ChineseCateringStatisticOpData -->
                    <div class="top1" style="overflow-y:hidden;overflow-x:hidden">
                        <e-charts :option="ChineseCateringStatisticOpData"/>
                    </div>
                    
                    <!-- ChineseCateringFundingStatisticTableData -->
                    <div class="down1">
                        <div id="scroll_in2" class="htp-list_scroll_in" style="height: 100%;overflow-y:hidden;overflow-x:hidden;font-size:20px;">
                            
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
                                <tr v-for="data in ChineseCateringFundingStatisticTableData">
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


                <div class="rightView">

                    <!-- whoAmIData -->
                    <div class="top0" style="overflow-y:hidden;overflow-x:hidden">    
                        <div class="table0Box1">
                            <div class="text0">
                                <h3>
                                    欢迎{{getIdentity()}} {{whoAmIData.RealName}}
                                </h3>
                            </div>
                        </div>
                        <div class="table0Box2">
                            <div class="text0">
                                Id : {{whoAmIData.Id}}
                                </div>
                        </div>
                        <div class="table0Box3">
                            <div class="text0">
                                用户名 : {{whoAmIData.Username}}
                            </div>
                        </div>
                    </div>

                    <!-- ChineseCateringPaymentTableData -->
                    <div class="top2" style="overflow-y:hidden;overflow-x:hidden">
                        <div class="table1" style="font-size:20px;">
                            <div class="table1Box" >
                                <div class="text" style="text-align: left;">2018年-2021年餐饮就业人数行业平均报酬</div>
                            </div>
                            <div class="table1Box">
                                <div class="table1BoxData1"><div class="text">时间</div></div>
                                <div class="table1BoxData2"><div class="text">平均报酬</div></div>
                                <div class="table1BoxData3"><div class="text">就业人数</div></div>
                            </div>
                            <div v-for="data in ChineseCateringPaymentTableData" class="table1Box">
                                <div class="table1BoxData1"><div class="text">{{data.Year}}</div></div>
                                <div class="table1BoxData2"><div class="text">{{data.AvgSalary}}</div></div>
                                <div class="table1BoxData3"><div class="text">{{data.CountSalary}}</div></div>
                            </div>
                        </div>
                        
                    </div>

                    <!-- ChineseCateringOnlineOrderStatisticOpData -->
                    <div class="down2" style="overflow-y:hidden;overflow-x:hidden">    
                        <e-charts :option="ChineseCateringOnlineOrderStatisticOpData"/>
                    </div>

                </div>
                    
            </div>

            <!-- Title -->
            <div class="Title" style="overflow-y:hidden;overflow-x:hidden">

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
    
    import * as PostApiUsers from '../../Web/PostApiUsers'
    import {postApiSelectChineseCateringStatistic} from '../../Web/PostApiSelectChineseCateringStatistic'
    import {postApiSelectChineseCateringOnlineOrderStatistic} from '../../Web/PostApiSelectChineseCateringOnlineOrderStatistic'
    import {postApiSelectChineseCateringPayment} from '../../Web/PostApiSelectChineseCateringPayment'
    import {postApiSelectChineseCateringBrandStatistic} from '../../Web/PostApiSelectChineseCateringBrandStatistic'
    import {postApiSelectChineseCateringFundingStatistic} from '../../Web/PostApiSelectChineseCateringFundingStatistic'
    export default{
        name: "login",

        data(){
            return{
                // 背景
                imgSrcBK:require('../../assets/background.jpeg'),
                ChineseCateringStatisticOpData : {},
                ChineseCateringOnlineOrderStatisticOpData : {},
                ChineseCateringPaymentTableData:{},
                ChineseCateringBrandStatisticOpData : {},
                ChineseCateringFundingStatisticTableData:{},
                whoAmIData:{},
            }
        },
        methods:{
            /**
             * @description: 用户类型参数转用户类型名
             * @return {*}
             */            
            getIdentity(){
                if (this.whoAmIData.Identity == '0') { return "管理员"}
                if (this.whoAmIData.Identity == '1') { return "投资方"}
                if (this.whoAmIData.Identity == '2') { return "从业者"}
                if (this.whoAmIData.Identity == '3') { return "游客"}
            },

            /**
             * @description: 注销登录
             * @return {*}
             */            
            logout(){
                this.$cookies.set("camp-session", "", -1)
                this.$router.push({
                    path: `/`
                })
            },

            /**
             * @description: 获取 ChineseCateringStatistic 的图表数据
             * @return {*}
             */            
            getChineseCateringStatisticOpData(){
                postApiSelectChineseCateringStatistic((res) =>{
                    this.ChineseCateringStatisticOpData = res.Data
                })
            },

            /**
             * @description: 获取 ChineseCateringOnlineOrderStatistic 的图表数据
             * @return {*}
             */            
            getChineseCateringOnlineOrderStatisticOpData(){
                postApiSelectChineseCateringOnlineOrderStatistic((res) =>{
                    this.ChineseCateringOnlineOrderStatisticOpData = res.Data
                })
            },

            /**
             * @description: 获取 ChineseCateringPayment 的表格数据
             * @return {*}
             */            
            getChineseCateringPaymentTableData(){
                postApiSelectChineseCateringPayment((res) =>{
                    this.ChineseCateringPaymentTableData = res.Data
                })
            },

            /**
             * @description: 获取 ChineseCateringBrandStatistic 的图表数据
             * @return {*}
             */            
            getChineseCateringBrandStatisticOpData(){
                postApiSelectChineseCateringBrandStatistic((res) =>{
                    console.log(res)
                    this.ChineseCateringBrandStatisticOpData = res.Data
                })
                
            },

            /**
             * @description: 获取 ChineseCateringFundingStatistic 的表格数据
             * @return {*}
             */            
            getChineseCateringFundingStatisticTableData(){
                postApiSelectChineseCateringFundingStatistic((res) =>{
                    this.ChineseCateringFundingStatisticTableData = res.Data
                })
            },

            /**
             * @description: 滚动显示控制
             * @param {*} srollId
             * @return {*}
             */            
            autoSroll(srollId) {
            // flag 为true时停止滚动
                var flag = false;
                // 定时器
                var timer;
                function roll() {
                    var h = -1;
                    timer = setInterval(function() {
                        flag = true;
                        if (flag) {
                            var current = document.getElementById(srollId).scrollTop;
                            if (current == h) {
                                //滚动到底端,返回顶端
                                h = 0;
                                document.getElementById(srollId).scrollTop = h + 1;
                            } else {
                                //以25ms/3.5px的速度滚动
                                h = current;
                                document.getElementById(srollId).scrollTop = h + 1;
                            }       
                        }
                    }, 50);
                }
                    
                document.getElementById(srollId).onclick = () => {       
                    if (flag) {
                        flag = false;
                        clearInterval(timer);
                    } else {
                        roll();
                    }
                };
                roll();
            },    
        },
        mounted(){
            PostApiUsers.postApiWhoAmI(this.$route.query.Id , (res)=>{
                this.whoAmIData = res.Data
            })
            this.getChineseCateringBrandStatisticOpData()
            this.getChineseCateringFundingStatisticTableData()
            this.getChineseCateringOnlineOrderStatisticOpData()
            this.getChineseCateringPaymentTableData()
            this.getChineseCateringStatisticOpData()
            this.autoSroll("scroll_in2")
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

