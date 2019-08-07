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

                <button type="button"
                    v-on:click="signin()">Sign In</button>
        </div>

        <!-- <p>{{ response }}</p> -->
    </div>

</template>

<script>
    import {HTTP} from '../http-constants';
    import axios from 'axios';


    export default {
        name: 'LogIn',
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
                        console.log(res);
                    })
                    .catch(e => {
                        this.errors = e
                    })
            }
        }
    }

</script>