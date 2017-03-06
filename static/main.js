// var chart1;
// var chart2;
// var chart3;

// var chart4;

function init() {
    var pDataChart1 = [{ x: [], y: [], type: "scatter" }];
    var pDataChart2 = [{ x: [], y: [], type: "scatter" }];
    var pDataChart3 = [{ x: [], y: [], type: "scatter" }];
    var pDataChart5 = [{ x: [], y: [], type: "scatter" }];
    var pDataChart4 = [{
        x: [],
        autobinx: false,
        histnorm: 'count',
        marker: { color: 'rgb(255, 217, 150)' },
        name: 'control',
        opacity: 0.75,
        type: 'histogram',
        xbins: {
            end: 1000,
            size: 1,
            start: -.5
        }
    }];
    var pDataChart6 = [{
        labels: ['<1ms', '1-5ms', '5-10ms', '10-50ms', '50-100ms', '>100ms'],
        type: 'pie',
        marker: { colors: ['#f7fcfd', '#e0ecf4', '#bfd3e6', '#9ebcda', '#8c96c6', '#8c6bb1', '#88419d', '#6e016b'] },
        sort: false,
        values: ['24', '21', '18', '15', '12', '2', '2', '6']
    }];


    Plotly.newPlot('c_gc', pDataChart1, {
        title: "GC Pauses",
        yaxis: {
            title: "Nanoseconds"
        }
    });
    Plotly.newPlot('c_mem', pDataChart2, {
        title: "Memory Allocated",
        yaxis: {
            title: "Bytes"
        }
    });
    Plotly.newPlot('c_thread', pDataChart3, {
        title: "Threads Num",
        yaxis: {
            title: "Num"
        }
    });
    Plotly.newPlot('c_goroutine', pDataChart5, {
        title: "Goroutine Num",
        yaxis: {
            title: "Num"
        }
    });


    Plotly.newPlot('c_latency_histo', pDataChart4, {
        bargap: 0.25,
        bargroupgap: 0.3,
        // barmode: 'overlay',
        title: 'SOIP Lantency Histogram',
        xaxis: { title: 'Latency(ms)' },
        yaxis: { title: 'Count' }
    });

    Plotly.newPlot(c_latency_pie, pDataChart6, {
        title: 'SOIP Lantency Pie_Chart',
        // xaxis: { title: 'Latency(ms)' },
        // yaxis: { title: 'Count' }
    });

}

init();


function wsurl() {
    var l = window.location;
    return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + "/debug/charts/data-feed";
}

ws = new WebSocket(wsurl());
ws.onopen = function() {
    ws.onmessage = function(evt) {
        var data = JSON.parse(evt.data);
        var t = moment(data.Time).format('HH:mm:ss')
        switch (data.Name) {
            case "thread":

                Plotly.extendTraces('c_thread', {
                    x: [
                        [t]
                    ],
                    y: [
                        [data.Value]
                    ]
                }, [0], 100);
                break;
            case "goroutine":
                Plotly.extendTraces('c_goroutine', {
                    x: [
                        [t]
                    ],
                    y: [
                        [data.Value]
                    ]
                }, [0], 100);
                break;
            case "mem":
                Plotly.extendTraces('c_mem', {
                    x: [
                        [t]
                    ],
                    y: [
                        [data.Value]
                    ]
                }, [0], 100);
                break;
            case "gc":
                Plotly.extendTraces('c_gc', {
                    x: [
                        [t]
                    ],
                    y: [
                        [data.Value]
                    ]
                }, [0], 100);
                break;
            case "latency":
                latencyPush(data.Value);
                break;
        }
    }
};


var latency_histo_delta = []
var latency_pie_values = [0, 0, 0, 0, 0, 0]


function latencyPush(latency) {
    latency_histo_delta.push(latency)
    if (latency < 1) {
        latency_pie_values[0]++
    } else if (1 <= latency && latency < 5) {
        latency_pie_values[1]++
    } else if (5 <= latency && latency < 10) {
        latency_pie_values[2]++
    } else if (10 <= latency && latency < 50) {
        latency_pie_values[3]++
    } else if (50 <= latency && latency < 100) {
        latency_pie_values[4]++
    } else {
        latency_pie_values[5]++
    }
}

if (ws != undefined) {
    var latency_interval = setInterval(function() {
        if (latency_histo_delta.length > 0)  {
            Plotly.extendTraces('c_latency_histo', {
                x: [
                    latency_histo_delta
                ]
            }, [0], -1);
            latencys = []
        }
        c_latency_pie.data[0].values = latency_pie_values
        Plotly.redraw(c_latency_pie);
    }, 1000);
}
