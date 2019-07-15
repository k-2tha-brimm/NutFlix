<template>
    <div class="movie">
        <h1>{{movie.title}}</h1>
    </div>
</template>

<script>
    import {HTTP} from '../http-constants';
    export default {
        name: 'Movie',
        props: ['movie'],
        data() {
            return {
                errors: ''
            }
        },
        // method: {
        //     getMovie: function (id) {
        //         HTTP.get(`/api/movies/${id}`)
        //         .then(res => {
        //             this.movie = res.data
        //         })
        //     }
        // }
    }
</script>
