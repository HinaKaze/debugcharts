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


    Plotly.newPlot('c_latency', pDataChart4, {
        bargap: 0.25,
        bargroupgap: 0.3,
        // barmode: 'overlay',
        title: 'SOIP Lantency',
        xaxis: { title: 'Latency(ms)' },
        yaxis: { title: 'Count' }
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
                }, [0], 86400);
                break;
            case "goroutine":
                Plotly.extendTraces('c_goroutine', {
                    x: [
                        [t]
                    ],
                    y: [
                        [data.Value]
                    ]
                }, [0], 86400);
                break;
            case "mem":
                Plotly.extendTraces('c_mem', {
                    x: [
                        [t]
                    ],
                    y: [
                        [data.Value]
                    ]
                }, [0], 86400);
                break;
            case "gc":
                Plotly.extendTraces('c_gc', {
                    x: [
                        [t]
                    ],
                    y: [
                        [data.Value]
                    ]
                }, [0], 86400);
                break;
            case "latency":
                Plotly.extendTraces('c_latency', {
                    x: [
                        [data.Value]
                    ]
                }, [0], 86400);
                break;
        }


        // var d = moment(data.Ts).format('HH:mm:ss');
        // if (data.GcPause != 0) {
        //     Plotly.extendTraces('container1', {
        //         x: [
        //             [d]
        //         ],
        //         y: [
        //             [data.GcPause]
        //         ]
        //     }, [0], 86400);
        // }
        // Plotly.extendTraces('container2', {
        //     x: [
        //         [d]
        //     ],
        //     y: [
        //         [data.BytesAllocated]
        //     ]
        // }, [0], 86400);
        // Plotly.extendTraces('container3', {
        //     x: [
        //         [d],
        //         [d]
        //     ],
        //     y: [
        //         [data.CpuSys],
        //         [data.CpuUser]
        //     ]
        // }, [0, 1], 86400);
    }
};
