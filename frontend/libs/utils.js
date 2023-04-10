import {useEffect, useState} from "react";
import {del} from "./api-admin";
import {toast} from "react-toastify";

export const useLocalStorage = (key, initialValue) => {
    const isServer = typeof window === 'undefined';
    // State to store our value
    // Pass initial state function to useState so logic is only executed on client
    const [storedValue, setStoredValue] = useState(() => {
        if (isServer) {
            return initialValue;
        }
        // Get from local storage by key
        const item = window.localStorage.getItem(key);
        // Parse stored json or if none return initialValue
        return item ? JSON.parse(item) : initialValue;
    });

    // useEffect to update local storage when component unmounts
    useEffect(() => {
        if (isServer) {
            return;
        }
        // Save state
        window.localStorage.setItem(key, JSON.stringify(storedValue));
    }, [storedValue]);

    return [storedValue, setStoredValue];
}
export const objToParams = (obj) => {
    if (!obj) {
        return
    }
    return Object.entries(obj).map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(v)}`).join('&');
}
export const handleChange = (setQuery, query, e) => {
    let v = e.target.value
    let name = e.target.name
    setQuery({...query, [name]: v});
}
export const handleChangeToLocalStorage = (key, setQuery, query, e) => {
    let v = e.target.value
    let name = e.target.name
    let newVar = {...query, [name]: v};
    setQuery(newVar);
    localStorage.setItem(key, JSON.stringify(newVar))
}
export const handleKeyPress = (setQuery, tempQuery, e) => {
    if (e.key === 'Enter') {
        setQuery(tempQuery)
    }
}
export const handleDel = async (url, id, mutate) => {
    if (confirm('确认删除?')) {
        const {code} = await del(`${url}?id=${id}`)
        if (code === 0) {
            toast.success('删除成功', {position: 'top-center'})
            mutate()
        }
    }
}
export const formatTags = (input) => {
    const tags = ['tag-info', 'tag-success', 'tag-primary', 'tag-warning', 'tag-brown', 'tag-purple', 'tag-danger']
    const values = input.split(',')
    let output = []
    for (let i = 0; i < values.length; i++) {
        output.push(`${values[i]}:${values[i]}:${tags[i % tags.length]}`);
    }
    return output.join(',')
}

export const makeUrl = (url, query) => {
    let q = objToParams(query)
    return `${url}?${q ? q : ''}`
}
export const redirect = (res, url,) => {
    res.writeHead(302, {Location: url || '/'})
    res.end()
}
export const dateTime = (inDate) => {
    const t = new Date(inDate);
    const now = new Date();
    const timeDiff = now - t;
    const seconds = timeDiff / 1000;
    const minutes = seconds / 60;
    const hours = minutes / 60;
    const days = hours / 24;
    switch (true) {
        case seconds < 60:
            return "几秒前";
        case minutes < 60:
            return Math.floor(minutes) + " 分钟前";
        case hours < 24:
            let date = new Date(new Date().getTime() - timeDiff);
            return `${date.getHours()} 小时 ${date.getMinutes()} 分钟前`;
        case days < 356:
            return Math.floor(days) + " 天前";
        default:
            return t.toString();
    }
}

export const getRandomNumber = (min, max) => {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

export const fetcher = (...args) => fetch(...args).then((res) => res.json())