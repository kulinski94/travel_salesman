// editor.js

new Vue({
  el: '#editor',
  data: {
    question: "",
    answer1: "",
    answer2: "",
    answer3: ""
  },
  methods: {
    save: function () {
      this.$http.post('api/question', { text: this.question, answers: [{text:this.answer1}, {text:this.answer2}, {text:this.answer3}] }).then((response) => {
        console.log(reponse);
      }, (error) => {
        console.error(error);
      });
    }
  }
})