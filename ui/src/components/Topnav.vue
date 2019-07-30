<template>
    <ul>
        <!-- <button @click='showModal = true'>Login</button> -->
        <router-link class="login" to="/login">Login</router-link>
        <modal v-if="showModal" @close="showModal = false"></modal>
    </ul>
</template>


<script>
import Modal from './Modal';
import LogInComponent from './Login.vue';

export default {
    name: 'TopNav',
    components: {
        'modal': Modal,
        'login': LogInComponent
    },
    data() {
        return {
            showModal: false
        }
    }
}
</script>

