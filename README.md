1. Get All Ads

Endpoint:
GET /ads

Description:
Fetches the list of all available ads.

Request:
No request body required.

Response:
Returns a JSON array of ads.

[
{
"id": 1,
"title": "Ad Title",
"description": "Ad Description",
"url": "https://example.com/ad1.mp4"
},
...
]

Status Codes:

    200 OK – On success

    500 Internal Server Error – If something goes wrong

2. Record Ad Click

Endpoint:
POST /ads/click

Description:
Records a click event for a given ad.

Request Body:

{
"ad_id": 1,
"timestamp": "2025-04-18T14:52:23Z",
"ip": "192.168.19.1",
"playback_position": 15.2,
"user_id": 12491
}

Response on Success:

{
"message": "Click recorded successfully"
}

Status Codes:

    200 OK – On success

    400 Bad Request – If input is invalid

    500 Internal Server Error – If storing in DB or Redis fails

3. Get Ad Analytics

Endpoint:
GET /ads/analytics

Description:
Fetches real-time or near real-time analytics data like Click Through Rate (CTR) or total Clicks for a given ad, based on a time window.

Request Body:

{
"ad_id": 1,
"type": "ctr",
"duration": "day"
}

Status Codes:

    200 OK – On success

    400 Bad Request – If parameters are invalid

    500 Internal Server Error – If a database error occurs
