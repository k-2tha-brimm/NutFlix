<template>
    <div class="movie">
        <h1>This is a test</h1>
    </div>
</template>

<script>
    import {HTTP} from '../http-constants';
    export default {
        name: 'Movie',
        data() {
            return {
                movie: {},
                errors: ''
            }
        },
        method: {
            getMovie: function (id) {
                HTTP.get(`/api/movies/${id}`)
                .then(res => {
                    this.movie = res.data
                })
            }
        }
    }
</script>
