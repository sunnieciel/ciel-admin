import {Footer, Headers, Nav, PageInfoWithInfo, searchDate, searchInput} from "../../../compoents/sys-page";
import Head from "next/head";
import {InfoBox} from "../../../compoents/toy";
import {useState} from "react";
import useSWR from "swr";
import {objToParams} from "../../../libs/utils";
import {service} from "../../../libs/request";


const pageConf = {
    infoBoxes: [
        '/image/p013.png:PayPal充值:t3:info-box mr-12',
        '/image/p001.png:支付宝充值:t1:info-box-primary mr-12',
        '/image/p003.png:微信充值:t2:info-box-success mr-12',
    ]
}
export default function WalletReport() {
    const [query, setQuery] = useState({uname: '', begin: '', end: ''})
    const [tempQuery, setTempQuery] = useState({})
    const {data} = useSWR(`/wallet/report?${objToParams(query)}`, service.get)
    return <>
        <Headers/>
        <div className={'wrapper '}>
            <div className={'w'}>
                <div className={'wrapper-left'}>
                    <Head>
                        <title>{process.env.SYSTEM_NAME} › 报表</title>
                    </Head>
                    <Nav/>
                    <PageInfoWithInfo backUrl={process.env.HOME_PAGE} backName={'BLEACH'} pageName={'报表'} icon={'/image/p001.png'}>
                        <div className={'cell flex'}>
                            <span className={'mr-auto'}></span>
                            {searchInput(1, '用户名', 'uname', tempQuery, setTempQuery, setQuery)}
                            {searchDate(2, tempQuery, setTempQuery, setQuery)}
                        </div>
                    </PageInfoWithInfo>
                    <div className={'box-02 no-bottom-border'}>
                        <>
                            <div className={'cell '}>
                                <span className={'color-desc-02 fs-13'}>默认统计最近半年</span>
                            </div>
                            <div className={'cell'}>
                                {pageConf.infoBoxes.map(i => {
                                    const arr = i.split(":")
                                    return <InfoBox img={arr[0]} title={arr[1]} num={data && data.data[arr[2]]} clas={arr[3]}/>
                                })}
                            </div>
                        </>
                        <div className={'cell-tools p-6'}></div>
                    </div>
                </div>
            </div>
        </div>
        <Footer/>
    </>
}