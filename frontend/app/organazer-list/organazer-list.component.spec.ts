import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrganazerListComponent } from './organazer-list.component';

describe('OrganazerListComponent', () => {
  let component: OrganazerListComponent;
  let fixture: ComponentFixture<OrganazerListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [OrganazerListComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OrganazerListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
