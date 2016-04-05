#include <grp.h>
#include <stdio.h>
#include <stdlib.h>
#include "_cgo_export.h"

void cEntryPoint() {
  struct group grp = {};
  char *buf = malloc(256);

  passingGrpBuffer(&grp, buf, 256);
  printf("grp.gr_name:   %s\n", grp.gr_name);
  printf("grp.gr_passwd: %s\n", grp.gr_passwd);
  printf("grp.gr_gid:    %d\n", grp.gr_gid);

  char *m;
  int idx;
  for (idx = 0; grp.gr_mem[idx] != NULL; idx++) {
    printf("  Member: %s\n", grp.gr_mem[idx]);
  }
}
