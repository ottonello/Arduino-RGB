int rpin = 11;
int gpin = 12;
int bpin = 13;

void setup() {
  pinMode(rpin, OUTPUT);
  pinMode(gpin, OUTPUT);
  pinMode(bpin, OUTPUT);

  Serial.begin(9600);
}

void loop() {
  int r, g, b = 0;
  while (Serial.available()) {
    byte c = Serial.read();
    if (c == 's') {
      Serial.print("Received ");

      while (!Serial.available()) {}
      r = Serial.read();

      Serial.print("r=");
      Serial.print(r);

      while (!Serial.available()) {}
      g = Serial.read();

      Serial.print(", g=");
      Serial.print(g);

      while (!Serial.available()) {}
      b = Serial.read();

      Serial.print(", b=");
      Serial.print(b);
      Serial.println();

      setColor(
        r, g, b
      );
    } else {
      Serial.print("Ignored ");
      Serial.print(c);
    }
  }
}

void setColor(int red, int green, int blue) {
  analogWrite(rpin, red);
  analogWrite(gpin, green);
  analogWrite(bpin, blue);
}
