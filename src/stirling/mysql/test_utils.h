#pragma once
#include <deque>
#include <string>
#include <utility>
#include "src/common/base/base.h"
#include "src/common/base/status.h"
#include "src/stirling/mysql/mysql.h"

namespace pl {
namespace stirling {
namespace mysql {
namespace testutils {

std::string GenRawPacket(uint8_t packet_num, std::string_view msg);

std::string GenRequest(char command, std::string_view msg);

Packet GenCountPacket(int num_col);

Packet GenColDefinition(const ColDefinition& col_def);

Packet GenResultsetRow(const ResultsetRow& row);

Packet GenStmtPrepareRespHeader(const StmtPrepareRespHeader& header);

Packet GenStmtExecuteRequest(const StmtExecuteRequest& req);

Packet GenStmtCloseRequest(const StmtCloseRequest& req);

Packet GenStringRequest(const StringRequest& req, MySQLEventType type);

std::deque<Packet> GenResultset(const Resultset& resultset, bool client_eof_deprecate = false);

std::deque<Packet> GenStmtPrepareOKResponse(const StmtPrepareOKResponse& resp);

Packet GenErr(const ErrResponse& err);

Packet GenOK();

Packet GenEOF();

}  // namespace testutils
}  // namespace mysql
}  // namespace stirling
}  // namespace pl
