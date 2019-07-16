<template>
    <ul>
        <button @click='showModal = true'>Login</button>
        <button @click='showModal = true'>Signup</button>
        <modal v-if="showModal" @close="showModal = false"></modal>
    </ul>
</template>


<script>
import Modal from './Modal';

export default {
    name: 'TopNav',
    components: {
        'modal': Modal
    },
    data() {
        return {
            showModal: false
        }
    }
}
</script>

