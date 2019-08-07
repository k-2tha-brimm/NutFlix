<template>
    <div class="login-form">
        <h2>Sign In</h2>
        <div class="form-group">
            <label for="email">
                Email:
            </label>
            <input type="text" 
                email="email"
                class="form-control"
                v-model="input.email"
                placeholder="Email">

            <label for="password">
                Password:
            </label>
            <input type="password"
                password="password"
                class="form-control"
                v-model="input.password"
                placeholder="Password">
                <router-link class="browse" to="/browse">Browse</router-link>
                    <button type="button"
                        v-on:click="signin()">Sign In</button>
        </div>

        <!-- <p>{{ response }}</p> -->
    </div>

</template>

<script>
    import {HTTP} from '../http-constants';
    import axios from 'axios';
    import Browse from './Browse.vue';


    export default {
        name: 'LogIn',
        components: {
            'browse': Browse
        },
        data() {
            return {
                input: {
                    email: '',
                    password: ''
                },
                errors: '',
                repsonse: ''
            }
        },
        methods: {
            signin: function () {
                
                let data = ({
                    email: this.input.email,
                    password: this.input.password
                });

                axios.post("http://localhost:5000/login", data, {
                    headers: {
                        'Content-Type': 'application/json'
                    }
                }).then(res => {
                    localStorage.setItem('user', JSON.stringify(res.data.email));
                    localStorage.setItem('session', true);
                    console.log(res);
                }).catch(e => {
                    this.errors = e
                })
            }
        }
    }

</script>