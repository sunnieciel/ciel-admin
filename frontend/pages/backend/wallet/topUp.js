import useSWR from "swr";
import {useState} from "react";
import {getChangeTypeOptions, listData, updateTopUpByAdmin} from "../../../libs/api-admin";
import {handleDel, objToParams, redirect,} from "../../../libs/utils";
import {FullScreenLoading} from "../../../compoents/common";
import {AddPage, Footer, Headers, Nav, PageButtons, PageInfo, SearchInput, Td, UpdatePage,} from "../../../compoents/sys-page";
import {toast} from "react-toastify";

const pageConf = {
    name: '充值订单', path: '/topUp',
    fields: [
        {field: 'id', name: 'Id', renderFn: (d) => d.id},
        {field: 'uname', name: '用户名', search: 1, required: 1, editHide: 1},
        {field: 'transId', name: '交易ID', search: 1, disabled: 1},
        {field: 'ip', name: 'IP', search: 1, disabled: 1},
        {field: 'changeType', name: '充值类型', search: 1, type: 'select', options: process.env.OPTIONS_STATUS, disabled: 1},
        {field: 'status', name: '状态', search: 1, options: '1:等待:tag-warning,2:成功:tag-success,3:失败:tag-danger', disabled: 1},
        {field: 'desc', name: '备注', search: 1},
        {field: 'aid', name: 'AID', search: 1, disabled: 1},
        {field: 'createdAt', name: '创建时间'},
    ]
}

export default function TopUp({options}) {
    const [query, setQuery] = useState() // 查询参数
    const [showType, setShowType] = useState(1) // 1 主页 2添加 3修改
    const [id, setId] = useState() // 修改数据时使用
    pageConf.fields[3].options = options
    return (<>
        <Headers/>
        <div className={'wrapper'}>
            <div className="w">
                <div className={'wrapper-left'}>
                    <Nav/>
                    {showType === 1 && <TopUpPage query={query} setQuery={setQuery} setShowType={setShowType} setId={setId}/>}
                    {showType === 2 && <AddPage pageConf={pageConf} setShowType={setShowType}/>}
                    {showType === 3 && <UpdatePage pageConf={pageConf} setShowType={setShowType} id={id}/>}
                </div>
            </div>
        </div>
        <Footer/>
    </>)
}

const TopUpPage = ({query, setQuery, setShowType, setId}) => {
    const [tempQuery, setTempQuery] = useState({})
    const s = objToParams(query)
    const {data, isLoading, mutate, error} = useSWR(`/${pageConf.path}/list?${s !== undefined ? s : ''}`, listData)
    const handleTopUp = async (type, id) => {
        const res = await updateTopUpByAdmin(type, id)
        if (res.code === 0) {
            toast.success('OK')
            mutate()
        }
    }
    if (error) return
    return <>
        <PageInfo>
            <div className={'cell p-3 flex-center'}>
                <span className={'mr-auto'}></span>
                <SearchInput pageConf={pageConf} tempQuery={tempQuery} setTempQuery={setTempQuery} setQuery={setQuery}></SearchInput>
            </div>
        </PageInfo>
        <div className={'box-02 no-bottom-border'}>
            {isLoading
                ? <FullScreenLoading/>
                : <>
                    {data && data.list.length === 0
                        ? <div className={'cell color-desc-02 fs-13'}>暂无数据</div>
                        : <>
                            <div className={'cell flex-center p-3'}><PageButtons query={query} setTempQuery={setTempQuery} setQuery={setQuery} totalPage={data && data.total_page} total={data && data.total}/></div>
                            <table className={'table-1'}>
                                <tbody>
                                <tr>{pageConf.fields.filter(i => !i.hide).map((i, index) => <th key={index}>{i.name}</th>)}
                                    <th>操作</th>
                                </tr>
                                {data && data.list.map((i, index) => <tr key={index}><Td pageConf={pageConf} data={i}/>
                                    <td>
                                        {i.status === 1
                                            ? <>
                                                <button className={'btn-success mr-6'} onClick={() => handleTopUp(1, i.id)}>通过</button>
                                                <button className={'btn-warning mr-6'} onClick={() => handleTopUp(2, i.id)}>拒绝</button>
                                            </>
                                            : <>
                                                <button className={'btn-warning mr-6'} onClick={() => setId(i.id) & setShowType(3)}>修改</button>
                                            </>
                                        }
                                        <button className={'btn-danger'} onClick={() => handleDel(pageConf.path, i.id, mutate)}>删除</button>
                                    </td>
                                </tr>)}
                                </tbody>
                            </table>
                            <div className={'cell-tools p-3 flex-center'}><PageButtons totalPage={data && data.total_page} total={data && data.total} query={query} setTempQuery={setTempQuery} setQuery={setQuery}/></div>
                        </>
                    }
                </>
            }
        </div>
    </>
}

export async function getServerSideProps({req, res}) {
    const {data, code} = await getChangeTypeOptions(req.cookies.admin_token)
    if (code !== 0) {
        redirect(res)
    }
    let arr = []
    data.forEach(i => arr.push(`${i.id}:${i.title}:${i.class}`))
    return {
        props: {
            options: arr.join(',')
        }
    }
}