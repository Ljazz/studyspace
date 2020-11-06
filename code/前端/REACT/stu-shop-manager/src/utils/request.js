import axios from 'axios';
import {getToken} from './auth';

const APIInstance = axios.create({
    // baseURL: 'http://localhost:3000',
    baseURL: 'http://47.93.11.106:4000/',
    timeout: 5000
});

// 全局请求拦截，发送请求之前执行
APIInstance.interceptors.request.use(function (config) {
    config.headers['authorization'] = 'Bearer' + getToken();
    return config;
}, function (error) {
    return Promise.reject(error);
});

// 请求之后执行
APIInstance.interceptors.response.use(function (response) {
    return response;
},function (error) {
    return Promise.reject(error);
});

export function get(url, params){
    return axios.get(url, {
        params
    })
}

export function post(url, data){
    return axios.post(url, data)
}