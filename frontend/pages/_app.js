import 'react-toastify/dist/ReactToastify.css';
import {toast, ToastContainer} from "react-toastify";
import React, {useEffect, useRef, useState} from "react";
import jsCookie from "js-cookie";
import {keyToken} from "../consts/consts";
import {useRouter} from "next/router";
import 'animate.css/animate.min.css';
import 'atropos/css'


export default function App({Component, pageProps}) {
    const [ws, setWs] = useState()
    const isMounted = useRef(false)
    const router = useRouter()
    const connectWs = () => {
        let token = jsCookie.get(keyToken)
        if (!ws && token) {
            const ws = new WebSocket(`${process.env.WS_API}?token=${token}`,)
            setWs(ws)
            ws.onopen = (e) => {
                console.log('SUCCEED')
            }
            ws.onerror = (e) => {
                // console.log('ERROR')
            }
            ws.onmessage = ({data}) => {
                const d = JSON.parse(data)
                toast(d.msg, {type: d.type, position: d.position ? d.position : 'bottom-right', autoClose: 10000, hideProgressBar: false})
                let utterance = new SpeechSynthesisUtterance(d.msg);
                speechSynthesis.speak(utterance);
            }
        }
    }
    useEffect(() => {
        if (isMounted.current) {
            return
        }
        isMounted.current = true
        if (router.pathname.startsWith('/backend')) {
            connectWs()
        }
    }, [])
    return (<>
        <Component  {...pageProps} ws={ws} connectWs={connectWs}/>
        <ToastContainer
            position={'top-center'}
            autoClose={3000}
            theme={'dark'}
            // hideProgressBar
        />
    </>);
}



