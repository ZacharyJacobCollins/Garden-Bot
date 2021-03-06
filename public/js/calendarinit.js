	$(document).ready(function() {


		/* initialize the external events
		-----------------------------------------------------------------*/

		$('#external-events .fc-event').each(function() {
			// store data so the calendar knows to render an event upon drop
			$(this).data('event', {
				title: $.trim($(this).text()), // use the element's text as the event title
				stick: true, // maintain when user navigates (see docs on the renderEvent method)
				start  : '2010-01-05',
                end    : '2010-01-07'
			});

			// make the event draggable using jQuery UI
			$(this).draggable({
				zIndex: 999,
				revert: true,      // will cause the event to go back to its
				revertDuration: 0  //  original position after the drag
			});
		});


		//Recurring events
		$('#external-events .recurring-event').each(function() {

			// store data so the calendar knows to render an event upon drop
			$(this).data('event', {
				title: $.trim($(this).text()), // use the element's text as the event title
				stick: true, // maintain when user navigates (see docs on the renderEvent method)
				dow: [ 1, 2, 3, 4 ] // Monday - Thursday
			});

			// make the event draggable using jQuery UI
			$(this).draggable({
				zIndex: 999,
				revert: true,      // will cause the event to go back to its
				revertDuration: 0  //  original position after the drag
			});
		});


		/* initialize the calendar
		-----------------------------------------------------------------*/
		$('#calendar').fullCalendar({
			header: {
				left: 'prev,next today',
				center: 'title',
				right: 'month,agendaWeekthisthis,agendaDay'
			},
			editable: true,
			droppable: true, // this allows things to be dropped onto the calendar
			drop: function () {
				console.log(this.innerHTML)
				console.log('adding event')
				$.ajax({
				    type: "POST",
				    url: "/addevent",
				    // The key needs to match your method's input parameter (case-sensitive).
				    data: JSON.stringify({ Type: this.innerHTML, StartTime: "8am", EndTime: "9am"}),
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
		});
	});
