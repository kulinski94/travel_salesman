// app.js

new Vue({
    // We want to target the div with an id of 'events'
    el: '#app',
    mounted: function () {
    },
    data: {
        cities: [],
        distance: 0,
        count: 5,
    },
    methods: {
        getPath: function () {
            this.$http.post('api/path', this.cities).then((response) => {
                var paths = JSON.parse(JSON.stringify(response.body.paths));
                this.distance = JSON.parse(JSON.stringify(response.body.distance));
                this.drawPath(paths);
            }, (error) => {
                console.error(error);
            });
        },
        getCities: function () {
            this.$http.get(`api/cities/${this.count}`).then((response) => {
                this.cities = JSON.parse(JSON.stringify(response.body));
                this.drawCities()
            }, (error) => {
                console.error(error);
            });
        },
        drawCities: function () {
            var elem = document.getElementById('map');
            var ctx = elem.getContext('2d');
            var height = elem.height;
            var width = elem.width;
            ctx.clearRect(0, 0, elem.width, elem.height)

            console.log(this.cities);
            this.cities.forEach(function (city) {
                ctx.fillStyle = "#FF0000";
                ctx.beginPath();
                ctx.arc(city.xCord, city.yCord, 3, 0, Math.PI * 2, true);
                ctx.closePath();
                ctx.fill();
            });
        },
        drawPath: function (paths) {
            var elem = document.getElementById('map');
            var ctx = elem.getContext('2d');
            var height = elem.height;
            var width = elem.width;

            ctx.strokeStyle = "#80D080";
            ctx.beginPath();
            // move to first node
            var n = paths[0];
            ctx.moveTo(n.xCord, n.yCord);

            for (var i = 0; i < paths.length; i++) {
                var city = paths[i];
                // // each node as a small dot
                var centerX = city.xCord;
                var centerY = height - city.yCord;
                ctx.fillStyle = "#308020";
                ctx.fillRect(centerX, centerY, 1, 1);
                // draw a line to the next node
                var nextCity;
                if (i + 1 == paths.length) {
                    nextCity = paths[0];
                } else {
                    nextCity = paths[i + 1];
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