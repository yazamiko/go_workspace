$(document).ready(function(e) {

	$( "#saveBtn" ).on( "click", function() {
		var files = document.getElementById("filePicker").files;
		var f = files[0];

		var reader = new FileReader();

		reader.onloadend = function(theFile) {
			if(theFile.target.readyState == FileReader.DONE) {
				var data = {
					title: $("#myName").val(),
					file: reader.result
				}
				$.post('save', data);
			}
		};

		reader.readAsBinaryString(f);
	});

});