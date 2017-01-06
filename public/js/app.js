var counter = 0;

var vm = new Vue({
	el: '#app',
	data: {

	}, 
	mounted: function() {
		console.log('vue instance mounted correctly')
	},
	methods: {
		addEvent() {
			console.log('adding event')
			$.ajax({
			    type: "POST",
			    url: "/addevent",
			    // The key needs to match your method's input parameter (case-sensitive).
			    data: JSON.stringify({ Type: "light", Id: counter}),
			    contentType: "application/json; charset=utf-8",
			    dataType: "json",
			    success: function(data){
			    	console.log('hit')
			    },
			    failure: function(errMsg) {
			    	console.log(err)
			    	console.log(errMsg)
			    }
			});
		},
	}
});






