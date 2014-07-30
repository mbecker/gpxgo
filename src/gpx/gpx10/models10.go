package gpx11

import (
	"encoding/xml"
)

/*

The GPX XML hierarchy:

gpx
    - attr: version (xsd:string) required
    - attr: creator (xsd:string) required
    name
    desc
    author
    email
    url
    urlname
    time
    keywords
    bounds
    wpt
        - attr: lat (gpx:latitudeType) required
        - attr: lon (gpx:longitudeType) required
        ele
        time
        magvar
        geoidheight
        name
        cmt
        desc
        src
        url
        urlname
        sym
        type
        fix
        sat
        hdop
        vdop
        pdop
        ageofdgpsdata
        dgpsid
    rte
        name
        cmt
        desc
        src
        url
        urlname
        number
        rtept
            - attr: lat (gpx:latitudeType) required
            - attr: lon (gpx:longitudeType) required
            ele
            time
            magvar
            geoidheight
            name
            cmt
            desc
            src
            url
            urlname
            sym
            type
            fix
            sat
            hdop
            vdop
            pdop
            ageofdgpsdata
            dgpsid
    trk
        name
        cmt
        desc
        src
        url
        urlname
        number
        trkseg
            trkpt
                - attr: lat (gpx:latitudeType) required
                - attr: lon (gpx:longitudeType) required
                ele
                time
                course
                speed
                magvar
                geoidheight
                name
                cmt
                desc
                src
                url
                urlname
                sym
                type
                fix
                sat
                hdop
                vdop
                pdop
                ageofdgpsdata
                dgpsid
*/