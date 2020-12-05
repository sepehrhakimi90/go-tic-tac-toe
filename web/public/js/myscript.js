var resetButton = document.querySelector('#restartButton');

var cells = document.getElementsByTagName('td');

//TODO Get player from server

function clearBoard() {
    for (var i = 0; i < cells.length; i++) {
        cells[i].textContent = '';
    }

}

$("#startButton").click(function (event) {
    event.preventDefault();
    var data = {}
    var firstPlayerName = $("#firstPlayerName").val();
    if(firstPlayerName==="") {
        firstPlayerName="player-1";
    }
    data["FirstPlayer"] = firstPlayerName;
    var secondPlayerName = $("#secondPlayerName").val();
    if(secondPlayerName==="") {
        secondPlayerName="player-2";
    }
    data["SecondPlayer"]=secondPlayerName;

    //ToDo Handel fail
    $.post({
        url: "newGame",
        data: data
        }).done(function (data) {
            if($("#startButton").text()==="Restart") {
                clearBoard()
                deactivateBoard()
            }
            if($(".card").length) {
                $(".card").remove()
            }
            addGameState(data["Data"]["Players"]);
            $("#gameBoard").removeClass("invisible");
            // $("#gameForm").addClass("invisible");
            // $("#restartButton").removeClass("invisible");
            $("#startButton").text("Restart")
            attachTableListener()
        })

    $("#gameBoard").show();
})


function addGameState(Players) {
    console.log(Players)
    var area = $("div#gameArea");
    $("div#board").remove();
    var element = `
    <div id="board" class="list-group col-lg-3 my-4">
                <li type="button" class="list-group-item list-group-item-action active">
                    ${Players[0].name}
                    <span class="badge badge-success">${Players[0].sign}</span>
                </li>

                <li type="button" class="list-group-item list-group-item-action">
                    ${Players[1].name}
                    <span class="badge badge-success">${Players[1].sign}</span>
                </li>
            </div>
    `
    area.prepend(element)
}

function attachTableListener() {
    $("#gameBoard td").each(function (idx) {
        $(this).click({"idx": idx},function (event) {
            cellClick(event)
        })
    })
}

function deactivateBoard() {
    $("#gameBoard td").each(function (idx) {
        $(this).off("click")
    })
}

function drawWinner(Player) {
    var area = $("div#gameArea");

    var element = `
        <div class="card col-lg-3">
            <div class="card-body">
                <button type="button" class="btn btn-success">!!!${Player.name} win!!!</button>
            </div>
        </div>
    `
    area.append(element)
}

function cellClick(event) {
    if(event.target.textContent !== ""){
        return
    }
    var idx = event.data.idx
    $.post({
        url: "move",
        data: {"selectedCell":idx}
    }).done(function (data) {
        var winner = data["data"]["winner"]
        var availableMoves = data["data"]["available_moves"]
        var current_player = data["data"]["current_player"]
        var prev_player = (current_player + 1) % 2

        if(winner != -1){
            console.log(data["data"]["Players"][winner]["name"] + "win")
            event.target.textContent =  data["data"]["Players"][prev_player]["sign"]
            deactivateBoard()
            drawWinner(data["data"]["Players"][winner])
        } else if (availableMoves <= 0){
            deactivateBoard()
            drawDraw()
        }
        var lis = $("#board li")
        $("#board li.active").removeClass("active")
        $(lis[current_player]).addClass("active")
        event.target.textContent =  data["data"]["Players"][prev_player]["sign"]
    })
}

function drawDraw() {
    var area = $("div#gameArea");

    var element = `
        <div class="card col-lg-3">
            <div class="card-body">
                <button type="button" class="btn btn-success">!!!DRAW!!!</button>
            </div>
        </div>
    `
    area.append(element)
}




