let imgArray = [
    "static/img/rock.png",
    "static/img/papel.png",
    "static/img/tijeras.png",
    "static/img/lagarto.png",
    "static/img/spock.png"
];

function choose(x) {
    fetch("/play?c=" + x)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            return response.json();
        })
        .then(data => {

            if ( x == 0){
                document.getElementById("player_choice").innerHTML = "El jugador eligío PIEDRA";
            }
            if ( x == 1){
                document.getElementById("player_choice").innerHTML = "El jugador eligío PAPEL";
            }
            if ( x == 2){
                document.getElementById("player_choice").innerHTML = "El jugador eligío TIJERA";
            }
            if ( x == 3){
                document.getElementById("player_choice").innerHTML = "El jugador eligío LAGARTO";
            }
            if ( x == 4){
                document.getElementById("player_choice").innerHTML = "El jugador eligío SPOCK";
            }

            document.getElementById("player_score").innerHTML = data.player_score;
            document.getElementById("computer_score").innerHTML = data.computer_score;
            
            document.getElementById("computer_choice").innerHTML = data.computer_choice;
            document.getElementById("round_result").innerHTML = data.round_result;
            document.getElementById("round_message").innerHTML = data.message;

            //Cargar imagenes
            var imgElement = document.getElementById("img_computer");
            imgElement.src = imgArray[data.computer_choice_int];
            
            

           // console.log(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}
