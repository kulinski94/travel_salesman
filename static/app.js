// app.js

new Vue({
    // We want to target the div with an id of 'events'
    el: '#app',
    mounted: function () {
        // When the application loads, we want to call the method that initializes
        // some data
        this.getCities();
    },
    data: {
        cities: []
    },
    methods: {
        getCities: function () {
            this.$http.get('api/cities').then((response) => {
                this.cities = JSON.parse(JSON.stringify(response.body));
            }, (error) => {
                console.error(error);
            });
        },
        findPath: function () {

            var elem = document.getElementById('map');
            var ctx = elem.getContext('2d');
            var height = elem.height;
            var width = elem.width;

            console.log(this.cities);
            this.cities.forEach(function (city) {
                ctx.fillStyle = "#FF0000";
                ctx.beginPath();
                ctx.arc(city.xCord, city.yCord, 3, 0, Math.PI * 2, true);
                ctx.closePath();
                ctx.fill();
            });

            ctx.strokeStyle = "#80D080";
            ctx.beginPath();
            // move to first node
            var n = this.cities[0];
            ctx.moveTo(n.xCord, n.yCord);

            for (var i = 0; i < this.cities.length; i++) {
                var city = this.cities[i];
                // // each node as a small dot
                var centerX = city.xCord;
                var centerY = height - city.yCord;
                ctx.fillStyle = "#208020";
                ctx.fillRect(centerX, centerY, 1, 1);
                // draw a line to the next node
                var nextCity;
                if (i + 1 == this.cities.length) {
                    nextCity = this.cities[0];
                }else{
                    nextCity = this.cities[i + 1];
                }
                
                ctx.lineTo(nextCity.xCord, nextCity.yCord);
            }
            // draw all lines
            ctx.lineWidth = 1;
            ctx.stroke();
            ctx.closePath();
        }
    }
});