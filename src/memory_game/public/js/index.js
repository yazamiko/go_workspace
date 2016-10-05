var selected = 0;
var score = 120;
var firstCardDiv;
var secondCardDiv;
var scoreInterval;
var cardTimeout;
var found = 0;
var manaPoint = 8;
var finalScore = 0;
const numOfCards = 6;

var cardsArray = ['abomination', 'squirrel', 'savannah', 'jack', 'hellscream', 
	'abomination', 'squirrel', 'savannah', 'jack', 'hellscream', 'bow', 'bow'];

String.prototype.filename=function(extension){
    var s= this.replace(/\\/g, '/');
    s= s.substring(s.lastIndexOf('/')+ 1);

    return extension? s.replace(/[?#].+$/, ''): s.split('.')[0];
}

function shuffle(array) {
  var currentIndex = array.length, temporaryValue, randomIndex;

  while (0 !== currentIndex) {

    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;

    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }

  return array;
}

function setScoreboard() {
	$.get('/scoreboard', function(data) {
		$('#scoreboard').empty();
		console.log(data);
		for (var i = 0; i < data.length; i++) {
			var div = $("<div class='row'>").append((i+1)+") "+data[i].Nickname+": "+data[i].Value+" ["+data[i].Inserted+"]"+"\n");
			$('#scoreboard').append(div);
		}
	});
}

function prepareGame() {
	$('#startBtn').show();
	$('#stopBtn').hide();

	$('.front').prop('hidden', true);
	$('.back').prop('hidden', false);

	prepareCards();

	$('#startBtn').on('click', start);
	$('#stopBtn').on('click', stop);
}

function submitScore() {
	data = {
		score: finalScore,
		nickname: $('#nickname').val()
	}
	$.ajax({
		url: 'save',
		type: 'POST',
		data: data,
	})
	.done(function() {
		console.log("success");
		setScoreboard();
	})
	.fail(function() {
		console.log("error");
	})
	.always(function() {
		console.log("complete");
	});
	
}

function start() {
	$('#startBtn').hide();

	revealCards();

	$('#score').text("Score: "+score);
	$(".card").on('click', handlerCardPick)
}

function stop() {
	clearInterval(scoreInterval);

	$('.front').prop('hidden', true);
	$('.back').prop('hidden', false);

	prepareCards();

	$('.card').unbind('click', handlerCardPick);

	score = 120;
	manaPoint = 0;

	updateManaPoint(8);
	updateScore();

	$('#stopBtn').hide();
	$('#startBtn').show();

}

function updateManaPoint(i) {
	manaPoint +=i;
	$('#mana').text("Mana Crystals: "+manaPoint);
}
function updateScore() {
	$('#score').text("Score: "+score);
}

function showPlusTen() {
	$('#plusSpan').prop('hidden', function( i, val ) {
	  return !val;
	});

	updateScore();

	setTimeout(function() {
		$('#plusSpan').prop('hidden', function( i, val ) {
		  return !val;
		});
	}, 1000);
}

function showMinusThree() {
	$('#minusSpan').prop('hidden', function( i, val ) {
	  return !val;
	});

	updateScore();

	setTimeout(function() {
		$('#minusSpan').prop('hidden', function( i, val ) {
		  return !val;
		});
	}, 1000);
}
function revealCards() {
	turnCard($('.card'));

	setTimeout(function() {
		turnCard($('.card'));
		$('#stopBtn').show();
		scoreInterval = setInterval(function() {
			score -= 2	;
			updateScore();
		}, 1000);
	}, 1750);
}

function prepareCards() {
	cardsArray = shuffle(cardsArray);
	
	$('.front').children().each(function(index, el) {
		el.src = "public/img/"+cardsArray[index]+".gif"
	});

}

function getCardName(div) {
	return div.children('.front').children('img').attr('src').filename();
}

function turnCard(div) {
	div.children().prop('hidden', function( i, val ) {
	  return !val;
	});
}

function isRevealed(div) {
	return div.children('.front').prop('hidden');
}

function checkState() {
	if(found == numOfCards) {
		found = 0;
		clearInterval(scoreInterval);

		$('.front').prop('hidden', true);
		$('.back').prop('hidden', false);

		prepareCards();

		$('.card').unbind('click', handlerCardPick);

		$('#stopBtn').hide();
		$('#startBtn').show();
	}
}

function checkManaPoint() {
	if(manaPoint == 0) {
		finalScore = score
		
		clearInterval(scoreInterval);

		$('.front').prop('hidden', true);
		$('.back').prop('hidden', false);

		prepareCards();

		$('.card').unbind('click', handlerCardPick);

		score = 120;
		manaPoint = 0;

		updateManaPoint(8);

		$('#stopBtn').hide();
		$('#startBtn').show();

		$('#score').text("Game Over\n(Score: "+finalScore+")");
		$('#modalScore').text(finalScore);
		$('#myModal').modal();
	} else {
		cardTimeout = setTimeout(function() {
			turnCard(firstCardDiv);
			turnCard(secondCardDiv);
			selected = 0;
		}, 1000)
	}
}

function checkMatch() {
	var firstCardName = getCardName(firstCardDiv);
	var secondCardName = getCardName(secondCardDiv);

	if(firstCardName == secondCardName) {
		score +=10;
		found++;
		selected = 0;
		showPlusTen();
	} else {
		score -= 3;
		updateManaPoint(-1);
		checkManaPoint();
		showMinusThree();
	}
}

function handlerCardPick() {
	var div = $(this);

	if(selected < 2 && manaPoint > 0) {
		if(isRevealed(div)) {
			selected++;

			turnCard(div);

			if(selected == 2) { 
				secondCardDiv = div;
				checkMatch();
				checkState();
			} else {
				firstCardDiv = div;
			}
		}
	}
}

$(document).ready(function(e) {
	prepareGame();
	setScoreboard();

	$('#submitBtn').on('click', submitScore);
});