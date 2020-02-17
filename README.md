# machinebox

## Что сделать
- чтение видеопотока
- каждый кадр проверяем через openCV
- при обнаружении лица отправка кадра в facebox
- если лицо распознано то "действие"
- если не распознано "научить"

```
facebox := facebox.New("http://localhost:8080")
openCV := opencv.New()

stream := CreateStreamFromLocalDevice("/dev/video0")

eventEmitter := eventemitter.New()

for !stream.Ended() {
  frame := stream.ReadFrame()
  if !openCV.FaceDetected() {
    continue
  }
  faces := facebox.Recognize(frame)
  
  for _, face := range faces {
    if face.Detected {
      // сделать что-то
      eventEmitter.Emit(JSON{"event": "face_detected", "who": face.Person})
    } else {
      // сделать что-то
      // скорее необходимо обучить систему, а кадр сохраняем и помечаем как "неизвестные"
    }
  }
}
```
