package main

// From FasfF1 python
//    'session_data': 'SessionData.json',  # track + session status + lap count
//    'session_info': 'SessionInfo.json',  # more rnd
//    'archive_status': 'ArchiveStatus.json',  # rnd=1880327548
//    'heartbeat': 'Heartbeat.jsonStream',  # Probably time synchronization?
//    'audio_streams': 'AudioStreams.jsonStream',  # Link to audio commentary
//    'driver_list': 'DriverList.jsonStream',  # Driver info and line story
//    'extrapolated_clock': 'ExtrapolatedClock.jsonStream',  # Boolean
//    'race_control_messages': 'RaceControlMessages.json',  # Flags etc
//    'session_status': 'SessionStatus.jsonStream',  # Start and finish times
//    'team_radio': 'TeamRadio.jsonStream',  # Links to team radios
//    'timing_app_data': 'TimingAppData.jsonStream',  # Tyres and laps (juicy)
//    'timing_stats': 'TimingStats.jsonStream',  # 'Best times/speed' useless
//    'track_status': 'TrackStatus.jsonStream',  # SC, VSC and Yellow
//    'weather_data': 'WeatherData.jsonStream',  # Temp, wind and rain
//    'position': 'Position.z.jsonStream',  # Coordinates, not GPS? (.z)
//    'car_data': 'CarData.z.jsonStream',  # Telemetry channels (.z)
//    'content_streams': 'ContentStreams.jsonStream',  # Lap by lap feeds
//    'timing_data': 'TimingData.jsonStream',  # Gap to car ahead
//    'lap_count': 'LapCount.jsonStream',  # Lap counter
//    'championship_prediction': 'ChampionshipPrediction.jsonStream'  # Points

const (
	TopicCarData             = "CarData.z" // Telemetry channels (.z)
	TopicDriverList          = "DriverList"
	TopicExtrapolatedClock   = "ExtrapolatedClock"
	TopicHeartbeat           = "Heartbeat"
	TopicLapCount            = "LapCount"
	TopicPosition            = "Position.z"
	TopicRaceControlMessages = "RaceControlMessages"
	TopicSessionData         = "SessionData"
	TopicTimingAppData       = "TimingAppData"
	TopicTimingData          = "TimingData"
	TopicTimingStats         = "TimingStats"
	TopicTopThree            = "TopThree"
	TopicTrackStatus         = "TrackStatus"
	TopicWeatherData         = "WeatherData"
)
