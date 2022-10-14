<!--
 * @Date: 2022-10-09 16:49:59
 * @LastEditTime: 2022-10-14 16:58:26
 * @FilePath: /frontEnd/src/views/display.vue
 * @Description: 数据显示页
-->
<template>
    <div class = "viewDisplay">
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="content">

<!------------------------------------------------------------------------------>
<!------------------------------------ 品牌数据  -------------------------------->
<!------------------------------------------------------------------------------>
            <div class="brandStatisticView"><e-charts :option="brandStatisticOpData"/></div> 

<!------------------------------------------------------------------------------>
<!------------------------------------ 市场数据  -------------------------------->
<!------------------------------------------------------------------------------>
            <div class="statisticView"><e-charts :option="statisticOpData"/></div> 

<!------------------------------------------------------------------------------>
<!------------------------------------ 融资数据 --------------------------------->
<!------------------------------------------------------------------------------>
            
            <!------------------------ 融资标题 ---------------------------------->
            <div class="fundingStatisticTitle" style="display:flex; flex-direction: column; flex: 100%;" >
                <div style="width: 100%; flex: 100%; text-align: left; color: black;" ><h3>
                    2021年-2022年预制菜行业融资情况
                </h3></div>
                <div   style="width: 100%; flex: 100%;">
                    <div style="width: 20%; left: 0%; position: absolute;"><b>品牌</b></div>
                    <div style="width: 20%; left: 20%; position: absolute;"><b>时间</b></div>
                    <div style="width: 20%; left: 40%; position: absolute;"><b>轮次</b></div>
                    <div style="width: 40%; left: 60%; position: absolute;"><b>投资商</b></div>
                </div> 
            </div>
            <!------------------------ 融资数据 ---------------------------------->
            <div class="fundingStatisticView">
                <div id="scroll_in2" class="htp-list_scroll_in" style="height: 100%;overflow-y:hidden;overflow-x:hidden;display:flex; flex-direction: column; flex: 100%; ">   
                    <div v-for="data in fundingStatisticTableData"  style="width: 100%; flex: 100%;  margin-top: 2.5%; padding-bottom: 17.5%; background-color: aquamarine;">
                            <div style="height: 100%; width: 20%; left: 0%; position: absolute;">{{data.brand}}</div>
                            <div style="height: 100%; width: 20%; left: 20%; position: absolute;">{{data.time}}</div>
                            <div style="height: 100%; width: 20%; left: 40%; position: absolute;">{{data.fundingRound}}</div>
                            <div style="height: 100%; width: 40%; left: 60%; position: absolute;">{{data.investor}}</div>
                    </div>    

                </div>
            </div> 

<!------------------------------------------------------------------------------>
<!------------------------------------ 用户数据  -------------------------------->
<!------------------------------------------------------------------------------>
            <div class="whoAmIView">
                    <h4>欢迎{{whoAmITableData.identity}}{{whoAmITableData.realName}}</h4>
                    用户名 : {{whoAmITableData.username}}
            </div> 

<!------------------------------------------------------------------------------>
<!------------------------------------ 薪资数据  -------------------------------->
<!------------------------------------------------------------------------------>
            <div class="paymentView" style="display:flex; flex-direction: column; flex: 100%;" >
                <div style="width: 100%; flex: 100%; text-align: left; color: black;" >
                    <h3>2018-2021 平均薪资和就业人数</h3>
                </div>
                <div style="width: 100%; flex: 100%;">
                    <div style="width: 30%; left: 0%; position: absolute;"><b>年份</b></div>
                    <div style="width: 40%; left:30%; position: absolute;"><b>平均工资</b></div>
                    <div style="width: 30%; left:70%; position: absolute;"><b>人数</b></div>
                </div>
                <div class="paymentTr" v-for="data in paymentTableData" style="width: 100%; flex: 100%; ">
                    <div style="width: 30%; left: 0%; position: absolute;">{{data.year}}</div>
                    <div style="width: 40%; left:30%; position: absolute;">{{data.avgSalary}}</div>
                    <div style="width: 35%; left:70%; position: absolute;">{{data.counts}}</div>
                </div>
            </div>

<!------------------------------------------------------------------------------>
<!------------------------------------ 订单数据  -------------------------------->
<!------------------------------------------------------------------------------> 
            <div class="onlineOrderStatisticView"><e-charts :option="onlineOrderStatisticOpData"/></div> 
        </div>

        <div class="title" style="display: flex;flex-direction: column;">
            <div v-for="c in '中国餐饮大数据平台'" style="width: 100%;">
                <h2>{{c}}</h2>
            </div>
            <div style="width:100%;">
                <button @click="logout()">登出<br></button>
            </div>
        </div>
    </div>
</template>

<script>
    import * as userApi from '../api/user'
    import * as brandStatisticApi from '../api/brandStatistic'
    import * as fundingStatisticApi from '../api/fundingStatistic'
    import * as onlineOrderStatisticApi from '../api/onlineOrderStatistic'
    import * as paymentApi from '../api/payment'
    import * as statisticApi from '../api/statistic'

    export default{
        data () {
            return {
                imgSrcBK:require('../assets/background.jpeg'),
                id : "",
                whoAmITableData:{},
                statisticOpData:{},
                paymentTableData:{},
                brandStatisticOpData:{},
                onlineOrderStatisticOpData:{},
                fundingStatisticTableData:{}
            }
        },
        methods : {
            /**
             * @description: 从 cookie 中读取 session值 
                * 若读取成功，通知后端删除对应的 session值， 同时删除 cookie 中的 session 值，然后跳转登录页
                * 若读取失败，直接跳转登录页
             * @return {*}
             */            
            logout(){
                let SessionKey = this.$cookies.get("camp-session")
                if ( SessionKey ) {
                    userApi.postLoginBySessionKeyApi(SessionKey, {})
                }
                this.$cookies.set("camp-session", "", -1)
                this.$router.push({
                    path: `/`
                })
                alert("登出成功")
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
        mounted() {
            // 获取火锅品牌数据
            brandStatisticApi.getDataFromBrandStatisticApi((res) => {
                this.brandStatisticOpData = res.data
            })
            // 获取融资数据
            fundingStatisticApi.getAllDataByFromFundingStatisticApi((res) =>{
                this.fundingStatisticTableData = res.data
            })
            // 获取网上订单数据
            onlineOrderStatisticApi.postAmountByYearAndMonthFromOnlineOrderStatisticApi(2010, 2021, (res) => {
                this.onlineOrderStatisticOpData = res.data
            })
            // 获取平均薪资和人数
            paymentApi.postAvgSalaryAndCountsByYearFromPaymentApi(2018, 2021, (res) => {
                // console.log("payment : ", res)
                this.paymentTableData = res.data
            })
            // 获取市场份额数据
            statisticApi.postAmountAndPercentageByYearFromStatisticApi(2014, 2021, (res) => {
                // console.log("statistic : ", res)
                this.statisticOpData = res.data
            })
            // 获取用户数据
            userApi.postWhoAmIApi(this.$route.query.id, (res) => {
                this.whoAmITableData = res.data
            })
            // 设置自动滚动
            this.autoSroll("scroll_in2")
        }
    }
</script>

<style>
/* ========================================================= */
/* =======================底层界面=========================== */
/* ========================================================= */
/* 总体设置 */
.viewDisplay{
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
        position: absolute;
        z-index: -1;
    }

    /* 内容设置 */
    .content{
        height: 100%; top: 0%;
        width: 95%; left: 0%;
        position: absolute;
        z-index: 1;
        opacity: 0.85;
    }
        
    /* 标题设置 */
    .title{
        height: 100%; top: 0%;
        left: 95%; width: 5%;
        position: absolute;
        z-index: 1;
        opacity: 0.7;
        background-color: rgb(124, 94, 16);
    }

/* ========================================================= */
/* ========================数据页============================ */
/* ========================================================= */

/* =============================================== */
/* ===================火锅数据视图================== */
/* =============================================== */
.brandStatisticView{
    height: 100%; top: 0%;
    width: 30%; left: 0%;
    position: absolute;
    background-color: #FFFAF5;
}

/* =============================================== */
/* ===================融资数据视图================== */
/* =============================================== */
/* -------------------融资标题视图------------------ */
.fundingStatisticTitle{
    height: 12%; top: 38%;
    width: 30%; left: 70%;
    position: absolute;
    background-color: #FFE0D5;
}

/* -------------------融资数据视图------------------ */
.fundingStatisticView{
    height: 50%; top: 50%;
    width: 30%; left: 70%;
    position: absolute;
    background-color: #FFE0D5;
}

/* =============================================== */
/* ===================市场份额视图================== */
/* =============================================== */
.statisticView{
    height: 50%; top: 0%;
    width: 40%; left: 30%;
    position: absolute;
    background-color: #FFF5EF;
    
}

/* =============================================== */
/* ===================用户数据视图================== */
/* =============================================== */
.whoAmIView{
    height: 12%; top: 0;
    width: 30%; left: 70%;
    position: absolute;
    background-color: #FFEAE0;
    
}

/* =============================================== */
/* ===================薪资数据视图================== */
/* =============================================== */
.paymentView{
    height: 26%; top: 12%;
    width: 30%; left: 70%;
    position: absolute;
    background-color: #FFE5DA;
}

/* =============================================== */
/* ===================订单数据视图================== */
/* =============================================== */
.onlineOrderStatisticView{
    height: 50%; top: 50%;
    width: 40%; left: 30%;
    position: absolute;
    background-color: #FFEFE5;
}



/* ========================================================= */
/* ==========================其他============================ */
/* ========================================================= */

/* =============================================== */
/* ===================滚动列表视图================== */
/* =============================================== */
.htp-list_scroll_in{
    height: 100%;
    width: 100%;
    top:0%;
    position: absolute;
}

#button{
    width: 40%;
    background-color: bisque;
}

</style>