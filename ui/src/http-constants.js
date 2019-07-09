import axios from 'axios';

let baseURL = "http://localhost:3000";

export const HTTP = axios.create(
    {
        baseURL: baseURL
    });