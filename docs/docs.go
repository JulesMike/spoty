// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jules Michael"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "checks if server is running",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.CheckerResult"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/health.CheckerResult"
                        }
                    }
                }
            }
        },
        "/api/authenticate": {
            "get": {
                "description": "redirects user to spotify for authentication",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spoty"
                ],
                "summary": "Authentication",
                "responses": {
                    "302": {
                        "description": "redirection to spotify",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "already authenticated",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/api/callback": {
            "get": {
                "description": "spotify redirects to the this endpoint on success",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spoty"
                ],
                "summary": "Callback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code from spotify",
                        "name": "code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "state from spotify",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "authenticated successfully",
                        "schema": {
                            "$ref": "#/definitions/http.Success"
                        }
                    },
                    "403": {
                        "description": "could not retrieve token",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "404": {
                        "description": "could not retrieve current user",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/api/current": {
            "get": {
                "description": "returns information about the current playing track",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spoty"
                ],
                "summary": "Current Playing Track",
                "responses": {
                    "200": {
                        "description": "returns full track information",
                        "schema": {
                            "$ref": "#/definitions/spotify.FullTrack"
                        }
                    },
                    "401": {
                        "description": "not authenticated",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "404": {
                        "description": "no current playing track found",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/api/current/images": {
            "get": {
                "description": "returns the album images of the current playing track",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spoty"
                ],
                "summary": "Album Images of Current Playing Track",
                "responses": {
                    "200": {
                        "description": "returns album images",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/spoty.Image"
                            }
                        }
                    },
                    "401": {
                        "description": "not authenticated",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "404": {
                        "description": "no current playing track found",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "500": {
                        "description": "album images could not be processed",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/api/version": {
            "get": {
                "description": "checks the server's version",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/build.Info"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "build.Info": {
            "type": "object",
            "properties": {
                "dirty_build": {
                    "type": "boolean"
                },
                "last_commit": {
                    "type": "string"
                },
                "revision": {
                    "type": "string"
                }
            }
        },
        "color.RGBA": {
            "type": "object",
            "properties": {
                "r": {
                    "type": "integer"
                }
            }
        },
        "health.CheckResult": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "Error contains the check error message, if the check failed.",
                    "type": "string"
                },
                "status": {
                    "description": "Status is the availability status of a component.",
                    "type": "string"
                },
                "timestamp": {
                    "description": "Timestamp holds the time when the check was executed.",
                    "type": "string"
                }
            }
        },
        "health.CheckerResult": {
            "type": "object",
            "properties": {
                "details": {
                    "description": "Details contains health information for all checked components.",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/health.CheckResult"
                    }
                },
                "status": {
                    "description": "Status is the aggregated system availability status.",
                    "type": "string"
                }
            }
        },
        "http.Error": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "A human-readable explanation specific to this\noccurrence of the problem.",
                    "type": "string"
                },
                "extensions": {
                    "description": "Extensions contains additional data.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "any"
                    }
                },
                "instance": {
                    "description": "A URI reference that identifies the specific\noccurrence of the problem.  It may or may not yield further\ninformation if dereferenced.",
                    "type": "string"
                },
                "status": {
                    "description": "The HTTP status code ([RFC7231], Section 6)\ngenerated by the origin server for this occurrence of the problem.",
                    "type": "integer"
                },
                "title": {
                    "description": "A short, human-readable summary of the problem\ntype.  It SHOULD NOT change from occurrence to occurrence of the\nproblem, except for purposes of localization (e.g., using\nproactive content negotiation; see [RFC7231], Section 3.4).",
                    "type": "string"
                },
                "type": {
                    "description": "A URI reference [RFC3986] that identifies the\nproblem type.  This specification encourages that, when\ndereferenced, it provide human-readable documentation for the\nproblem type (e.g., using HTML [W3C.REC-html5-20141028]).  When\nthis member is not present, its value is assumed to be\n\"about:blank\".",
                    "type": "string"
                }
            }
        },
        "http.Success": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "spotify.FullTrack": {
            "type": "object",
            "properties": {
                "album": {
                    "description": "The album on which the track appears. The album object includes a link in href to full information about the album.",
                    "$ref": "#/definitions/spotify.SimpleAlbum"
                },
                "artists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spotify.SimpleArtist"
                    }
                },
                "available_markets": {
                    "description": "A list of the countries in which the track can be played,\nidentified by their ISO 3166-1 alpha-2 codes.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "disc_number": {
                    "description": "The disc number (usually 1 unless the album consists of more than one disc).",
                    "type": "integer"
                },
                "duration_ms": {
                    "description": "The length of the track, in milliseconds.",
                    "type": "integer"
                },
                "explicit": {
                    "description": "Whether or not the track has explicit lyrics.\ntrue =\u003e yes, it does; false =\u003e no, it does not.",
                    "type": "boolean"
                },
                "external_ids": {
                    "description": "Known external IDs for the track.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "external_urls": {
                    "description": "External URLs for this track.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "href": {
                    "description": "A link to the Web API endpoint providing full details for this track.",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_playable": {
                    "description": "IsPlayable defines if the track is playable. It's reported when the \"market\" parameter is passed to the tracks\nlisting API.\nSee: https://developer.spotify.com/documentation/general/guides/track-relinking-guide/",
                    "type": "boolean"
                },
                "linked_from": {
                    "description": "LinkedFrom points to the linked track. It's reported when the \"market\" parameter is passed to the tracks listing\nAPI.",
                    "$ref": "#/definitions/spotify.LinkedFromInfo"
                },
                "name": {
                    "type": "string"
                },
                "popularity": {
                    "description": "Popularity of the track.  The value will be between 0 and 100,\nwith 100 being the most popular.  The popularity is calculated from\nboth total plays and most recent plays.",
                    "type": "integer"
                },
                "preview_url": {
                    "description": "A URL to a 30 second preview (MP3) of the track.",
                    "type": "string"
                },
                "track_number": {
                    "description": "The number of the track.  If an album has several\ndiscs, the track number is the number on the specified\nDiscNumber.",
                    "type": "integer"
                },
                "type": {
                    "description": "Type of the track",
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "spotify.Image": {
            "type": "object",
            "properties": {
                "height": {
                    "description": "The image height, in pixels.",
                    "type": "integer"
                },
                "url": {
                    "description": "The source URL of the image.",
                    "type": "string"
                },
                "width": {
                    "description": "The image width, in pixels.",
                    "type": "integer"
                }
            }
        },
        "spotify.LinkedFromInfo": {
            "type": "object",
            "properties": {
                "external_urls": {
                    "description": "ExternalURLs are the known external APIs for this track or album",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "href": {
                    "description": "Href is a link to the Web API endpoint providing full details",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the linked track",
                    "type": "string"
                },
                "type": {
                    "description": "Type of the link: album of the track",
                    "type": "string"
                },
                "uri": {
                    "description": "URI is the Spotify URI of the track/album",
                    "type": "string"
                }
            }
        },
        "spotify.SimpleAlbum": {
            "type": "object",
            "properties": {
                "album_group": {
                    "description": "The field is present when getting an artist’s\nalbums. Possible values are “album”, “single”,\n“compilation”, “appears_on”. Compare to album_type\nthis field represents relationship between the artist\nand the album.",
                    "type": "string"
                },
                "album_type": {
                    "description": "The type of the album: one of \"album\",\n\"single\", or \"compilation\".",
                    "type": "string"
                },
                "artists": {
                    "description": "A slice of SimpleArtists",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spotify.SimpleArtist"
                    }
                },
                "available_markets": {
                    "description": "The markets in which the album is available,\nidentified using ISO 3166-1 alpha-2 country\ncodes.  Note that al album is considered\navailable in a market when at least 1 of its\ntracks is available in that market.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "external_urls": {
                    "description": "Known external URLs for this album.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "href": {
                    "description": "A link to the Web API enpoint providing full\ndetails of the album.",
                    "type": "string"
                },
                "id": {
                    "description": "The SpotifyID for the album.",
                    "type": "string"
                },
                "images": {
                    "description": "The cover art for the album in various sizes,\nwidest first.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spotify.Image"
                    }
                },
                "name": {
                    "description": "The name of the album.",
                    "type": "string"
                },
                "release_date": {
                    "description": "The date the album was first released.  For example, \"1981-12-15\".\nDepending on the ReleaseDatePrecision, it might be shown as\n\"1981\" or \"1981-12\". You can use ReleaseDateTime to convert this\nto a time.Time value.",
                    "type": "string"
                },
                "release_date_precision": {
                    "description": "The precision with which ReleaseDate value is known: \"year\", \"month\", or \"day\"",
                    "type": "string"
                },
                "uri": {
                    "description": "The SpotifyURI for the album.",
                    "type": "string"
                }
            }
        },
        "spotify.SimpleArtist": {
            "type": "object",
            "properties": {
                "external_urls": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "href": {
                    "description": "A link to the Web API enpoint providing full details of the artist.",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "uri": {
                    "description": "The Spotify URI for the artist.",
                    "type": "string"
                }
            }
        },
        "spoty.Image": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "hex": {
                    "type": "string"
                },
                "rgba": {
                    "$ref": "#/definitions/color.RGBA"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.2.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Spoty API",
	Description:      "Access information about current playing track on spotify through REST endpoints.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
