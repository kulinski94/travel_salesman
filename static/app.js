// app.js

new Vue({
    // We want to target the div with an id of 'events'
    el: '#app',
    mounted: function () {
        // When the application loads, we want to call the method that initializes
        // some data
        this.getQuestions();
    },
    data: {
        questions: [],
        picked: ""
    },
    methods: {
        getQuestions: function () {
            this.$http.get('api/question').then((response) => {
                this.questions = JSON.parse(JSON.stringify(response.body));
            }, (error) => {
                console.error(error);
            });
        },
        deleteQuestion: function(id) {
            this.$http.delete('api/question/' + id).then((response) => {
                console.log(response);
            }, (error) => {
                console.error(error);
            });
        },
        saveToPdf: function () {
            this.$http.get('api/quiz/pdf').then((response) => {
                console.log(response);
            }, (error) => {
                console.error(error);
            });
        }
    }
});