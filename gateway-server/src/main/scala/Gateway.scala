import py4j.GatewayServer

object Gateway {
  def main(args: Array[String]): Unit = {
    // Start a GatewayServer on an ephemeral port
    val gatewayServer: GatewayServer = new GatewayServer(null, 0)
    gatewayServer.start()
    val boundPort: Int = gatewayServer.getListeningPort
    if (boundPort == -1) {
      System.out.print("GatewayServer failed to bind; exiting")
      System.exit(1)
    } else {
      System.out.print(s"Started PythonGatewayServer on port $boundPort")
    }

    while (System.in.read() != -1) {}

    System.out.print("Exiting due to broken pipe from Python driver")
  }
}
