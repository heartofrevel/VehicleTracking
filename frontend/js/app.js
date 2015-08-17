var mapOptions = {
        center: new google.maps.LatLng(0, 0),
        zoom: 17,
        mapTypeId: google.maps.MapTypeId.ROADMAP
    };

var map = new google.maps.Map(document.getElementById('googleMap'), mapOptions);

function init() {	
	document.getElementById("myBtn").addEventListener("click", function(){
		setInterval(function(){
			$.ajax({
				type: "GET",
				url: "/tracker",
				cache: false,
				success: function(result){
					var jsonObj = jQuery.parseJSON(result);
					var lat = parseFloat(jsonObj.Latitude);
					var lon = parseFloat(jsonObj.Longitude);
					var myLatlng = new google.maps.LatLng(lat,lon);
					var center = new google.maps.LatLng(lat, lon);
					
					map.panTo(center);	
					var marker = new google.maps.Marker({
						position: myLatlng,
						map: map,
					});
				}
			});
		}, 5000);
	});
	
}
google.maps.event.addDomListener(window, 'load', init);